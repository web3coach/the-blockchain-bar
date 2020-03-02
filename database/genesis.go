package database

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
)

var genesisJson = `
{
  "genesis_time": "2019-03-18T00:00:00.000000000Z",
  "chain_id": "the-blockchain-bar-ledger",
  "balances": {
    "0x22ba1F80452E6220c7cc6ea2D1e3EEDDaC5F694A": 1000000
  }
}`

type Genesis struct {
	Balances map[common.Address]uint `json:"balances"`
}

func loadGenesis(path string) (Genesis, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return Genesis{}, err
	}

	var loadedGenesis Genesis
	err = json.Unmarshal(content, &loadedGenesis)
	if err != nil {
		return Genesis{}, err
	}

	return loadedGenesis, nil
}

func writeGenesisToDisk(path string, genesis []byte) error {
	return ioutil.WriteFile(path, genesis, 0644)
}
