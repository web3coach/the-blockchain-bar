package node

import (
	"context"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/web3coach/the-blockchain-bar/database"
	"github.com/web3coach/the-blockchain-bar/wallet"
	"testing"
	"time"
)

func TestValidBlockHash(t *testing.T) {
	hexHash := "000000fa04f8160395c387277f8b2f14837603383d33809a4db586086168edfa"
	var hash = database.Hash{}

	hex.Decode(hash[:], []byte(hexHash))

	isValid := database.IsBlockHashValid(hash)
	if !isValid {
		t.Fatalf("hash '%s' starting with 6 zeroes is suppose to be valid", hexHash)
	}
}

func TestInvalidBlockHash(t *testing.T) {
	hexHash := "000001fa04f8160395c387277f8b2f14837603383d33809a4db586086168edfa"
	var hash = database.Hash{}

	hex.Decode(hash[:], []byte(hexHash))

	isValid := database.IsBlockHashValid(hash)
	if isValid {
		t.Fatal("hash is not suppose to be valid")
	}
}

func TestMine(t *testing.T) {
	miner := database.NewAccount(wallet.AndrejAccount)
	pendingBlock := createRandomPendingBlock(miner)

	ctx := context.Background()

	minedBlock, err := Mine(ctx, pendingBlock)
	if err != nil {
		t.Fatal(err)
	}

	minedBlockHash, err := minedBlock.Hash()
	if err != nil {
		t.Fatal(err)
	}

	if !database.IsBlockHashValid(minedBlockHash) {
		t.Fatal()
	}

	if minedBlock.Header.Miner.String() != miner.String() {
		t.Fatal("mined block miner should equal miner from pending block")
	}
}

func TestMineWithTimeout(t *testing.T) {
	miner := database.NewAccount(wallet.AndrejAccount)
	pendingBlock := createRandomPendingBlock(miner)

	ctx, _ := context.WithTimeout(context.Background(), time.Microsecond*100)

	_, err := Mine(ctx, pendingBlock)
	if err == nil {
		t.Fatal(err)
	}
}

func createRandomPendingBlock(miner common.Address) PendingBlock {
	return NewPendingBlock(
		database.Hash{},
		0,
		miner,
		[]database.Tx{
			database.Tx{From: database.NewAccount(wallet.AndrejAccount), To: database.NewAccount(wallet.BabaYagaAccount), Value: 1, Time: 1579451695, Data: ""},
		},
	)
}
