package dto

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/web3coach/the-blockchain-bar/database"
)

type PendingTx struct {
	database.SignedTx
	Hash string `json:"hash"`
}

func (t PendingTx) MarshalJSON() ([]byte, error) {

	type ptx struct {
		From     common.Address `json:"from"`
		To       common.Address `json:"to"`
		Gas      uint           `json:"gas"`
		GasPrice uint           `json:"gasPrice"`
		Value    uint           `json:"value"`
		Nonce    uint           `json:"nonce"`
		Data     string         `json:"data"`
		Time     uint64         `json:"time"`
		Sig      []byte         `json:"signature"`
		Hash     string         `json:"hash"`
	}

	return json.Marshal(ptx{
		From:     t.From,
		To:       t.To,
		Gas:      t.Gas,
		GasPrice: t.GasPrice,
		Value:    t.Value,
		Nonce:    t.Nonce,
		Data:     t.Data,
		Time:     t.Time,
		Sig:      t.Sig,
		Hash:     t.Hash,
	})
}
