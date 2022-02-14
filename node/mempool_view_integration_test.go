package node

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/web3coach/the-blockchain-bar/database"
	"github.com/web3coach/the-blockchain-bar/fs"
	"github.com/web3coach/the-blockchain-bar/wallet"
)

func TestNode_MempoolViewer(t *testing.T) {

	babaYaga := database.NewAccount(testKsBabaYagaAccount)
	andrej := database.NewAccount(testKsAndrejAccount)

	// test cases
	poolLen := 3
	txn3From := babaYaga

	dataDir, err := getTestDataDirPath()
	if err != nil {
		t.Fatal(err)
	}

	genesisBalances := make(map[common.Address]uint)
	genesisBalances[andrej] = 1000000
	genesisBalances[babaYaga] = 1000000
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
	tx3 := database.NewBaseTx(babaYaga, andrej, 1, 1, "")

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

	signedTx3, err := wallet.SignTxWithKeystoreAccount(tx3, babaYaga, testKsAccountsPwd, wallet.GetKeystoreDirPath(dataDir))
	if err != nil {
		t.Error(err)
		return
	}

	// Add 3 new TXs
	err = n.AddPendingTX(signedTx1, nInfo)
	if err != nil {
		t.Fatal(err)
	}
	err = n.AddPendingTX(signedTx2, nInfo)
	if err != nil {
		t.Fatal(err)
	}
	err = n.AddPendingTX(signedTx3, nInfo)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/mempool/", nil)

	func(w http.ResponseWriter, r *http.Request, node *Node) {
		mempoolViewer(w, r, node.pendingTXs)
	}(rr, req, n)

	if rr.Code != http.StatusOK {
		t.Fatal("unexpected status code: ", rr.Code, rr.Body.String())
	}

	var resp map[string]database.SignedTx
	dec := json.NewDecoder(rr.Body)
	err = dec.Decode(&resp)
	if err != nil {
		t.Fatal("error decoding", err)
	}

	// check pool length
	if len(resp) != poolLen {
		t.Fatalf("mempool viewer reponse len wrong, got %v; want %v", len(resp), poolLen)
	}

	for _, v := range resp {
		// check for third case
		if v.From.Hex() == txn3From.Hex() {
			if !reflect.DeepEqual(signedTx3.Sig, v.Sig) {
				t.Errorf("invalid signature for txn, got %q, want %q", base64.StdEncoding.EncodeToString(v.Sig), base64.StdEncoding.EncodeToString(signedTx3.Sig))
			}
		}
	}
}
