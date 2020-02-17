package database

import (
	"crypto/sha256"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"time"
)

func NewAccount(value string) common.Address {
	return common.HexToAddress(value)
}

type Tx struct {
	From  common.Address `json:"from"`
	To    common.Address `json:"to"`
	Value uint           `json:"value"`
	Data  string         `json:"data"`
	Time  uint64         `json:"time"`
}

func NewTx(from, to common.Address, value uint, data string) Tx {
	return Tx{from, to, value, data, uint64(time.Now().Unix())}
}

func (t Tx) IsReward() bool {
	return t.Data == "reward"
}

func (t Tx) Hash() (Hash, error) {
	txJson, err := json.Marshal(t)
	if err != nil {
		return Hash{}, err
	}

	return sha256.Sum256(txJson), nil
}
