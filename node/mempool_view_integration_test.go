package node

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/web3coach/the-blockchain-bar/database"
	"github.com/web3coach/the-blockchain-bar/dto"
	"github.com/web3coach/the-blockchain-bar/fs"
	"github.com/web3coach/the-blockchain-bar/wallet"
)

func TestNode_MempoolViewer(t *testing.T) {

	babaYaga := database.NewAccount(testKsBabaYagaAccount)
	andrej := database.NewAccount(testKsAndrejAccount)

	// test cases
	poolLen := 2
	wantNonce := uint(2)
	wantAccount := andrej

	dataDir, err := getTestDataDirPath()
	if err != nil {
		t.Fatal(err)
	}

	genesisBalances := make(map[common.Address]uint)
	genesisBalances[andrej] = 1000000
	genesis := database.Genesis{Balances: genesisBalances, ForkTIP1: 0}
	genesisJson, err := json.Marshal(genesis)
	if err != nil {
		t.Fatal(err)
	}

	err = database.InitDataDirIfNotExists(dataDir, genesisJson)
	defer fs.RemoveDir(dataDir)

	err = copyKeystoreFilesIntoTestDataDirPath(dataDir)
	if err != nil {
		t.Fatal(err)
	}

	// Required for AddPendingTX() to describe
	// from what node the TX came from (local node in this case)
	nInfo := NewPeerNode(
		"127.0.0.1",
		8085,
		false,
		database.NewAccount(""),
		true,
		nodeVersion,
	)

	// Start mining with a high mining difficulty, just to be slow on purpose and let a synced block arrive first
	n := New(dataDir, nInfo.IP, nInfo.Port, babaYaga, nInfo, nodeVersion, uint(5))

	state, err := database.NewStateFromDisk(n.dataDir, n.miningDifficulty)
	if err != nil {
		t.Fatal(err)
	}
	defer state.Close()

	n.state = state

	pendingState := state.Copy()
	n.pendingState = &pendingState

	tx1 := database.NewBaseTx(andrej, babaYaga, 1, 1, "")
	tx2 := database.NewBaseTx(andrej, babaYaga, 2, 2, "")

	signedTx1, err := wallet.SignTxWithKeystoreAccount(tx1, andrej, testKsAccountsPwd, wallet.GetKeystoreDirPath(dataDir))
	if err != nil {
		t.Error(err)
		return
	}

	signedTx2, err := wallet.SignTxWithKeystoreAccount(tx2, andrej, testKsAccountsPwd, wallet.GetKeystoreDirPath(dataDir))
	if err != nil {
		t.Error(err)
		return
	}

	// Add 2 new TXs into the BabaYaga's node, triggers mining

	err = n.AddPendingTX(signedTx1, nInfo)
	if err != nil {
		t.Fatal(err)
	}
	err = n.AddPendingTX(signedTx2, nInfo)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/mempool/", nil)

	func(w http.ResponseWriter, r *http.Request, node *Node) {
		mempoolViewer(w, r, node.pendingTXs)
	}(rr, req, n)

	if rr.Code != http.StatusOK {
		t.Error("unexpected status code: ", rr.Code, rr.Body.String())
	}

	resp := []dto.PendingTx{}
	dec := json.NewDecoder(rr.Body)
	err = dec.Decode(&resp)
	if err != nil {
		t.Error("error decoding", err)
	}

	// sort by timestamp before test
	sort.SliceStable(resp, func(i, j int) bool {
		return resp[i].Nonce < resp[j].Nonce
	})

	// check pool length
	if len(resp) != poolLen {
		t.Errorf("mempool viewer reponse len wrong, got %v; want %v", len(resp), poolLen)
	}

	// first txn is from andrej
	gotAccount := resp[0].From
	if gotAccount != wantAccount {
		t.Errorf("txn from invalid, got %q, want %q", gotAccount, wantAccount)
	}

	// second nonce
	gotNonce := resp[1].Nonce
	if gotNonce != wantNonce {
		t.Errorf("mempool viewer unexpected nonce, got %v; want %v", gotNonce, wantNonce)
	}
}
