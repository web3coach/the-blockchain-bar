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
package node

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/web3coach/the-blockchain-bar/database"
	"github.com/web3coach/the-blockchain-bar/wallet"
	"net/http"
	"strconv"
)

type ErrRes struct {
	Error string `json:"error"`
}

type BalancesRes struct {
	Hash     database.Hash           `json:"block_hash"`
	Balances map[common.Address]uint `json:"balances"`
}

type TxAddReq struct {
	From    string `json:"from"`
	FromPwd string `json:"from_pwd"`
	To      string `json:"to"`
	Value   uint   `json:"value"`
	Data    string `json:"data"`
}

type TxAddRes struct {
	Success bool `json:"success"`
}

type StatusRes struct {
	Hash       database.Hash       `json:"block_hash"`
	Number     uint64              `json:"block_number"`
	KnownPeers map[string]PeerNode `json:"peers_known"`
	PendingTXs []database.SignedTx `json:"pending_txs"`
}

type SyncRes struct {
	Blocks []database.Block `json:"blocks"`
}

type AddPeerRes struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func listBalancesHandler(w http.ResponseWriter, r *http.Request, state *database.State) {
	enableCors(&w)

	writeRes(w, BalancesRes{state.LatestBlockHash(), state.Balances})
}

func txAddHandler(w http.ResponseWriter, r *http.Request, node *Node) {
	req := TxAddReq{}
	err := readReq(r, &req)
	if err != nil {
		writeErrRes(w, err)
		return
	}

	from := database.NewAccount(req.From)

	if from.String() == common.HexToAddress("").String() {
		writeErrRes(w, fmt.Errorf("%s is an invalid 'from' sender", from.String()))
		return
	}

	if req.FromPwd == "" {
		writeErrRes(w, fmt.Errorf("password to decrypt the %s account is required. 'from_pwd' is empty", from.String()))
		return
	}

	nonce := node.state.GetNextAccountNonce(from)
	tx := database.NewTx(from, database.NewAccount(req.To), req.Value, nonce, req.Data)

	signedTx, err := wallet.SignTxWithKeystoreAccount(tx, from, req.FromPwd, wallet.GetKeystoreDirPath(node.dataDir))
	if err != nil {
		writeErrRes(w, err)
		return
	}

	err = node.AddPendingTX(signedTx, node.info)
	if err != nil {
		writeErrRes(w, err)
		return
	}

	writeRes(w, TxAddRes{Success: true})
}

func statusHandler(w http.ResponseWriter, r *http.Request, node *Node) {
	enableCors(&w)

	res := StatusRes{
		Hash:       node.state.LatestBlockHash(),
		Number:     node.state.LatestBlock().Header.Number,
		KnownPeers: node.knownPeers,
		PendingTXs: node.getPendingTXsAsArray(),
	}

	writeRes(w, res)
}

func syncHandler(w http.ResponseWriter, r *http.Request, node *Node) {
	reqHash := r.URL.Query().Get(endpointSyncQueryKeyFromBlock)

	hash := database.Hash{}
	err := hash.UnmarshalText([]byte(reqHash))
	if err != nil {
		writeErrRes(w, err)
		return
	}

	blocks, err := database.GetBlocksAfter(hash, node.dataDir)
	if err != nil {
		writeErrRes(w, err)
		return
	}

	writeRes(w, SyncRes{Blocks: blocks})
}

func addPeerHandler(w http.ResponseWriter, r *http.Request, node *Node) {
	peerIP := r.URL.Query().Get(endpointAddPeerQueryKeyIP)
	peerPortRaw := r.URL.Query().Get(endpointAddPeerQueryKeyPort)
	minerRaw := r.URL.Query().Get(endpointAddPeerQueryKeyMiner)

	peerPort, err := strconv.ParseUint(peerPortRaw, 10, 32)
	if err != nil {
		writeRes(w, AddPeerRes{false, err.Error()})
		return
	}

	peer := NewPeerNode(peerIP, peerPort, false, database.NewAccount(minerRaw), true)

	node.AddPeer(peer)

	fmt.Printf("Peer '%s' was added into KnownPeers\n", peer.TcpAddress())

	writeRes(w, AddPeerRes{true, ""})
}
