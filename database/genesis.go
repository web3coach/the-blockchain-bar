package database

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
)

var genesisJson = `{
  "genesis_time": "2020-06-01T00:00:00.000000000Z",
  "chain_id": "the-blockchain-bar-ledger",
  "symbol": "TBB",
  "balances": {
    "0x09eE50f2F37FcBA1845dE6FE5C762E83E65E755c": 1000000
  }
}`

type Genesis struct {
	Balances map[common.Address]uint `json:"balances"`
	Symbol   string                  `json:"symbol"`
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
