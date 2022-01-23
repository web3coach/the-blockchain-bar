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
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"reflect"
)

func GetBlocksAfter(blockHash Hash, dataDir string) ([]Block, error) {
	f, err := os.OpenFile(getBlocksDbFilePath(dataDir), os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}

	blocks := make([]Block, 0)
	shouldStartCollecting := false

	if reflect.DeepEqual(blockHash, Hash{}) {
		shouldStartCollecting = true
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		var blockFs BlockFS
		err = json.Unmarshal(scanner.Bytes(), &blockFs)
		if err != nil {
			return nil, err
		}

		if shouldStartCollecting {
			blocks = append(blocks, blockFs.Value)
			continue
		}

		if blockHash == blockFs.Key {
			shouldStartCollecting = true
		}
	}

	return blocks, nil
}

// GetBlockByHeightOrHash returns the requested block by hash or height.
// Add a map cache of height and hash to State struct
// and the corresponding functions to update and read this info
// The value of the maps corresponds to the position on the file corresponding to the block

//
//  type State struct {
//    Balances      map[common.Address]uint
//    Account2Nonce map[common.Address]uint
//
//    HashCache     map[string]uint64
//    HeightCache   map[uint64]uint64
//
//    dbFile *os.File
//
//    latestBlock     Block
//    latestBlockHash Hash
//    hasGenesisBlock bool
//
//    miningDifficulty uint
//
//    forkTIP1 uint64
//  }
//
func GetBlockByHeightOrHash(state *State, height uint64, hash, dataDir string) (Block, Hash, error) {

	blk := Block{}
	hsh := Hash{}

	key, ok := state.HeightCache[height]
	if hash != "" {
		key, ok = state.HashCache[hash]
	}

	if !ok {
		return blk, hsh, errors.New("invalid block hash or height")
	}

	f, err := os.OpenFile(getBlocksDbFilePath(dataDir), os.O_RDONLY, 0600)
	if err != nil {
		return blk, hsh, err
	}
	defer f.Close()

	f.Seek(key, 0)
	scanner := bufio.NewScanner(f)
	if scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return blk, hsh, err
		}
		var blockFs BlockFS
		err = json.Unmarshal(scanner.Bytes(), &blockFs)
		if err != nil {
			return blk, hsh, err
		}
		blk = blockFs.Value
		hsh = blockFs.Key
	}

	return blk, hsh, nil
}
