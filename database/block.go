package database

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
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
	TXs    []Tx        `json:"payload"`
}

type BlockHeader struct {
	Parent Hash    `json:"parent"`
	Number uint64  `json:"number"`
	Nonce  uint32  `json:"nonce"`
	Time   uint64  `json:"time"`
	Miner  Account `json:"miner"`
}

type BlockFS struct {
	Key   Hash  `json:"hash"`
	Value Block `json:"block"`
}

func NewBlock(parent Hash, number uint64, nonce uint32, time uint64, miner Account, txs []Tx) Block {
	return Block{BlockHeader{parent, number, nonce, time, miner}, txs}
}

func (b Block) Hash() (Hash, error) {
	blockJson, err := json.Marshal(b)
	if err != nil {
		return Hash{}, err
	}

	return sha256.Sum256(blockJson), nil
}

func IsBlockHashValid(hash Hash) bool {
	return fmt.Sprintf("%x", hash[0]) == "0" &&
		fmt.Sprintf("%x", hash[1]) == "0" &&
		fmt.Sprintf("%x", hash[2]) == "0" &&
		fmt.Sprintf("%x", hash[3]) != "0"
}
