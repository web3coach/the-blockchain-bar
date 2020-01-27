package database

import (
	"crypto/sha256"
	"encoding/json"
	"time"
)

type Account string

func NewAccount(value string) Account {
	return Account(value)
}

type Tx struct {
	From  Account `json:"from"`
	To    Account `json:"to"`
	Value uint    `json:"value"`
	Data  string  `json:"data"`
	Time  uint64  `json:"time"`
}

func NewTx(from Account, to Account, value uint, data string) Tx {
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
