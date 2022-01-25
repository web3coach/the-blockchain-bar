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
	"fmt"
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
// It uses cached data in the State struct (HashCache / HeightCache)
func GetBlockByHeightOrHash(state *State, height uint64, hash, dataDir string) (Block, Hash, error) {

	blk := Block{}
	hsh := Hash{}

	key, ok := state.HeightCache[height]
	if hash != "" {
		key, ok = state.HashCache[hash]
	}

	if !ok {
		if hash != "" {
			return blk, hsh, fmt.Errorf("invalid hash: '%v'", hash)
		}
		return blk, hsh, fmt.Errorf("invalid height: '%v'", height)
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
