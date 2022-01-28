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
	"os"
	"reflect"
)

type BlockFSResult struct {
	BlockFS BlockFS
	Err     error
}

func GetBlocksAfterHash(blockHash Hash, dataDir string) ([]Block, error) {
	blocks := make([]Block, 0)

	shouldStartCollecting := false
	if reflect.DeepEqual(blockHash, Hash{}) {
		shouldStartCollecting = true
	}

	for val := range readBlockFS(dataDir) {
		if val.Err != nil {
			return nil, val.Err
		}

		if shouldStartCollecting {
			blocks = append(blocks, val.BlockFS.Value)
			continue
		}

		if blockHash == val.BlockFS.Key {
			shouldStartCollecting = true
		}
	}
	return nil, nil
}

func GetBlockByHeight(limit uint64, dataDir string) (*Block, error) {
	for val := range readBlockFS(dataDir) {
		if val.Err != nil {
			return nil, val.Err
		}
		if limit == val.BlockFS.Value.Header.Number {
			return &val.BlockFS.Value, nil
		}
	}
	return nil, nil
}

func GetBlockByHash(blockHash Hash, dataDir string) (*Block, error) {
	for val := range readBlockFS(dataDir) {
		if val.Err != nil {
			return nil, val.Err
		}
		if blockHash == val.BlockFS.Key {
			return &val.BlockFS.Value, nil
		}
	}
	return nil, nil
}

func readBlockFS(dataDir string) chan BlockFSResult {
	ch := make(chan BlockFSResult)

	f, err := os.OpenFile(getBlocksDbFilePath(dataDir), os.O_RDONLY, 0600)
	if err != nil {
		ch <- BlockFSResult{Err: err}
		return ch
	}

	go func() {
		defer f.Close()
		defer close(ch)

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if err := scanner.Err(); err != nil {
				ch <- BlockFSResult{Err: err}
				return
			}

			var blockFs BlockFS
			err = json.Unmarshal(scanner.Bytes(), &blockFs)
			if err != nil {
				ch <- BlockFSResult{Err: err}
				return
			}

			ch <- BlockFSResult{BlockFS: blockFs}
		}
	}()
	return ch
}
