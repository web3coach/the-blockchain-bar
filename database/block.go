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
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

const BlockReward = 100

type Hash [32]byte

func (h Hash) MarshalText() ([]byte, error) {
	return []byte(h.Hex()), nil
}

func (h *Hash) UnmarshalText(data []byte) error {
	_, err := hex.Decode(h[:], data)
	return err
}

func (h Hash) Hex() string {
	return hex.EncodeToString(h[:])
}

func (h Hash) IsEmpty() bool {
	emptyHash := Hash{}

	return bytes.Equal(emptyHash[:], h[:])
}

type Block struct {
	Header BlockHeader `json:"header"`
	TXs    []SignedTx  `json:"payload"`
}

type BlockHeader struct {
	Parent Hash           `json:"parent"`
	Number uint64         `json:"number"`
	Nonce  uint32         `json:"nonce"`
	Time   uint64         `json:"time"`
	Miner  common.Address `json:"miner"`
}

type BlockFS struct {
	Key   Hash  `json:"hash"`
	Value Block `json:"block"`
}

func NewBlock(parent Hash, number uint64, nonce uint32, time uint64, miner common.Address, txs []SignedTx) Block {
	return Block{BlockHeader{parent, number, nonce, time, miner}, txs}
}

func (b Block) Hash() (Hash, error) {
	blockJson, err := json.Marshal(b)
	if err != nil {
		return Hash{}, err
	}

	return sha256.Sum256(blockJson), nil
}

func (b Block) GasReward() uint {
	reward := uint(0)

	for _, tx := range b.TXs {
		reward += tx.GasCost()
	}

	return reward
}

func IsBlockHashValid(hash Hash, miningDifficulty uint) bool {
	zeroesCount := uint(0)

	for i := uint(0); i < miningDifficulty; i++ {
		if fmt.Sprintf("%x", hash[i]) == "0" {
			zeroesCount++
		}
	}

	if fmt.Sprintf("%x", hash[miningDifficulty]) == "0" {
		return false
	}

	return zeroesCount == miningDifficulty
}
