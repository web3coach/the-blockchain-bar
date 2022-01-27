package node

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/web3coach/the-blockchain-bar/database"
)

func TestBlockExplorer(t *testing.T) {

	tc := []struct {
		arg  string
		want uint64
	}{
		{"12", 12},
		{"34", 34},
		{"0000003dc60b50d8f98e5e49f1cf520a84f95f51890849b1ac37eda6c07718df", 8},
		{"000000244ab3ada6479fd06f0eb81b3b97051859191380758cc546bfe2074759", 2},
		{"99", 99}, // this must return http.Status != 200
	}
	datadir := "test_block_explorer_db"

	n := New(datadir, "127.0.0.1", 8085, database.NewAccount(DefaultMiner), PeerNode{}, nodeVersion, 3)

	t.Log(fmt.Sprintf("Listening on: %s:%d", n.info.IP, n.info.Port))

	state, err := database.NewStateFromDisk(n.dataDir, n.miningDifficulty)
	if err != nil {
		t.Fatal(err)
	}
	defer state.Close()

	n.state = state

	pendingState := state.Copy()
	n.pendingState = &pendingState

	t.Log("Blockchain state:")
	t.Logf("	- height: %d\n", n.state.LatestBlock().Header.Number)
	t.Logf("	- hash: %s\n", n.state.LatestBlockHash().Hex())

	for _, tc := range tc {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/block/"+tc.arg, nil)

		func(w http.ResponseWriter, r *http.Request, node *Node) {
			blockByNumberOrHash(w, r, node)
		}(rr, req, n)

		if rr.Code != http.StatusOK {
			if tc.want == 99 { // this is an error case, so continue
				continue
			}
			t.Error("unexpected status code: ", rr.Code, rr.Body.String())
		}

		resp := new(database.BlockFS)
		dec := json.NewDecoder(rr.Body)
		err = dec.Decode(resp)
		if err != nil {
			t.Error("error decoding", err)
		}

		got := resp.Value.Header.Number
		if got != tc.want {
			t.Errorf("block explorer(%q)  = %v; want %v", tc.arg, got, tc.want)
		}
	}
}
