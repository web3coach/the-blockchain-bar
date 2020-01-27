package node

import (
	"context"
	"github.com/web3coach/the-blockchain-bar/database"
	"github.com/web3coach/the-blockchain-bar/fs"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNode_Run(t *testing.T) {
	datadir := getTestDataDirPath()
	err := fs.RemoveDir(datadir)
	if err != nil {
		t.Fatal(err)
	}

	n := New(datadir, "127.0.0.1", 8085, PeerNode{})

	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	err = n.Run(ctx)
	if err.Error() != "http: Server closed" {
		t.Fatal("node server was suppose to close after 5s")
	}
}

func TestNode_Mining(t *testing.T) {
	datadir := getTestDataDirPath()
	err := fs.RemoveDir(datadir)
	if err != nil {
		t.Fatal(err)
	}

	n := New(datadir, "127.0.0.1", 8085, PeerNode{})
	ctx, closeNode := context.WithTimeout(context.Background(), time.Minute*15)

	go func() {
		time.Sleep(time.Second * 1)
		tx := database.NewTx("andrej", "babayaga", 1, "")
		myself := NewPeerNode("127.0.0.1", 8085, false, true)
		_ = n.AddPendingTX(tx, myself)
	}()

	go func() {
		time.Sleep(time.Second * 30)
		tx := database.NewTx("andrej", "babayaga", 2, "")
		myself := NewPeerNode("127.0.0.1", 8085, false, true)
		_ = n.AddPendingTX(tx, myself)
	}()

	go func() {
		ticker := time.NewTicker(10 * time.Second)

		for {
			select {
			case <-ticker.C:
				if n.state.LatestBlock().Header.Number == 2 {
					closeNode()
					return
				}
			}
		}
	}()

	_ = n.Run(ctx)

	if n.state.LatestBlock().Header.Number != 2 {
		t.Fatal("was suppose to mine 2 pending TX into 2 valid blocks under 30m")
	}
}

func TestNode_MiningStopsOnNewSyncedBlock(t *testing.T) {
	datadir := getTestDataDirPath()
	err := fs.RemoveDir(datadir)
	if err != nil {
		t.Fatal(err)
	}

	n := New(datadir, "127.0.0.1", 8085, PeerNode{})
	ctx, closeNode := context.WithTimeout(context.Background(), time.Minute*15)

	tx := database.Tx{From: "andrej", To: "babayaga", Value: 1, Time: 1579451695, Data: ""}
	tx2 := database.NewTx("andrej", "babayaga", 2, "")
	tx2Hash, _ := tx2.Hash()

	validSyncedBlock := database.NewBlock(database.Hash{}, 1, 1453450257, 1579451704, []database.Tx{tx})

	go func() {
		time.Sleep(time.Second * (miningIntervalSeconds - 2))

		myself := NewPeerNode("127.0.0.1", 8085, false, true)
		err := n.AddPendingTX(tx, myself)
		if err != nil {
			t.Fatal(err)
		}

		err = n.AddPendingTX(tx2, myself)
		if err != nil {
			t.Fatal(err)
		}
	}()

	go func() {
		time.Sleep(time.Second * (miningIntervalSeconds + 2))
		if !n.isMining {
			t.Fatal("should be mining")
		}

		_, err := n.state.AddBlock(validSyncedBlock)
		if err != nil {
			t.Fatal(err)
		}
		n.newSyncedBlocks <- validSyncedBlock

		time.Sleep(time.Second * 2)
		if n.isMining {
			t.Fatal("new received block should have canceled mining")
		}

		_, onlyTX2IsPending := n.pendingTXs[tx2Hash.Hex()]

		if len(n.pendingTXs) != 1 && !onlyTX2IsPending {
			t.Fatal("new received block should have canceled mining of already mined transaction")
		}

		time.Sleep(time.Second * (miningIntervalSeconds + 2))
		if !n.isMining {
			t.Fatal("should be mining again the 1 TX not included in synced block")
		}
	}()

	go func() {
		ticker := time.NewTicker(10 * time.Second)

		for {
			select {
			case <-ticker.C:
				if n.state.LatestBlock().Header.Number == 2 {
					closeNode()
					return
				}
			}
		}
	}()

	_ = n.Run(ctx)

	if n.state.LatestBlock().Header.Number != 2 {
		t.Fatal("was suppose to mine 2 pending TX into 2 valid blocks under 30m")
	}

	if len(n.pendingTXs) != 0 {
		t.Fatal("no pending TXs should be left to mine")
	}
}

func getTestDataDirPath() string {
	return filepath.Join(os.TempDir(), ".tbb_test")
}
