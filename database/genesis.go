// Copyright 2020 The the-blockchain-bar Authors
// This file is part of the the-blockchain-bar library.
//
// The the-blockchain-bar library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The the-blockchain-bar library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.
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
