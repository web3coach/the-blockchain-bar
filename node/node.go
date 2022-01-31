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
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/caddyserver/certmagic"
	"github.com/ethereum/go-ethereum/common"

	"github.com/web3coach/the-blockchain-bar/database"
)

const DefaultBootstrapIp = "node.tbb.web3.coach"

// The Web3Coach's Genesis account with 1M TBB tokens
const DefaultBootstrapAcc = "0x09ee50f2f37fcba1845de6fe5c762e83e65e755c"
const DefaultMiner = "0x0000000000000000000000000000000000000000"
const DefaultIP = "127.0.0.1"
const HttpSSLPort = 443
const endpointStatus = "/node/status"

const endpointSync = "/node/sync"
const endpointSyncQueryKeyFromBlock = "fromBlock"

const endpointAddPeer = "/node/peer"
const endpointAddPeerQueryKeyIP = "ip"
const endpointAddPeerQueryKeyPort = "port"
const endpointAddPeerQueryKeyMiner = "miner"
const endpointAddPeerQueryKeyVersion = "version"

const endpointBlockByNumberOrHash = "/block/"

const miningIntervalSeconds = 10
const DefaultMiningDifficulty = 3

type PeerNode struct {
	IP          string         `json:"ip"`
	Port        uint64         `json:"port"`
	IsBootstrap bool           `json:"is_bootstrap"`
	Account     common.Address `json:"account"`
	NodeVersion string         `json:"node_version"`

	// Whenever my node already established connection, sync with this Peer
	connected bool
}

func (pn PeerNode) TcpAddress() string {
	return fmt.Sprintf("%s:%d", pn.IP, pn.Port)
}

func (pn PeerNode) ApiProtocol() string {
	if pn.Port == HttpSSLPort {
		return "https"
	}

	return "http"
}

type Node struct {
	dataDir string
	info    PeerNode

	// The main blockchain state after all TXs from mined blocks were applied
	state *database.State

	// temporary pending state validating new incoming TXs but reset after the block is mined
	pendingState *database.State

	knownPeers      map[string]PeerNode
	pendingTXs      map[string]database.SignedTx
	archivedTXs     map[string]database.SignedTx
	newSyncedBlocks chan database.Block
	newPendingTXs   chan database.SignedTx
	nodeVersion     string

	// Number of zeroes the hash must start with to be considered valid. Default 3
	miningDifficulty uint
	isMining         bool
}

func New(dataDir string, ip string, port uint64, acc common.Address, bootstrap PeerNode, version string, miningDifficulty uint) *Node {
	knownPeers := make(map[string]PeerNode)

	n := &Node{
		dataDir:          dataDir,
		info:             NewPeerNode(ip, port, false, acc, true, version),
		knownPeers:       knownPeers,
		pendingTXs:       make(map[string]database.SignedTx),
		archivedTXs:      make(map[string]database.SignedTx),
		newSyncedBlocks:  make(chan database.Block),
		newPendingTXs:    make(chan database.SignedTx, 10000),
		nodeVersion:      version,
		isMining:         false,
		miningDifficulty: miningDifficulty,
	}

	n.AddPeer(bootstrap)

	return n
}

func NewPeerNode(ip string, port uint64, isBootstrap bool, acc common.Address, connected bool, version string) PeerNode {
	return PeerNode{ip, port, isBootstrap, acc, version, connected}
}

func (n *Node) Run(ctx context.Context, isSSLDisabled bool, sslEmail string) error {
	fmt.Println(fmt.Sprintf("Listening on: %s:%d", n.info.IP, n.info.Port))

	state, err := database.NewStateFromDisk(n.dataDir, n.miningDifficulty)
	if err != nil {
		return err
	}
	defer state.Close()

	n.state = state

	pendingState := state.Copy()
	n.pendingState = &pendingState

	fmt.Println("Blockchain state:")
	fmt.Printf("	- height: %d\n", n.state.LatestBlock().Header.Number)
	fmt.Printf("	- hash: %s\n", n.state.LatestBlockHash().Hex())

	go n.sync(ctx)
	go n.mine(ctx)

	return n.serveHttp(ctx, isSSLDisabled, sslEmail)
}

func (n *Node) LatestBlockHash() database.Hash {
	return n.state.LatestBlockHash()
}

func (n *Node) serveHttp(ctx context.Context, isSSLDisabled bool, sslEmail string) error {
	handler := http.NewServeMux()

	handler.HandleFunc("/balances/list", func(w http.ResponseWriter, r *http.Request) {
		listBalancesHandler(w, r, n.state)
	})

	handler.HandleFunc("/tx/add", func(w http.ResponseWriter, r *http.Request) {
		txAddHandler(w, r, n)
	})

	handler.HandleFunc(endpointStatus, func(w http.ResponseWriter, r *http.Request) {
		statusHandler(w, r, n)
	})

	handler.HandleFunc(endpointSync, func(w http.ResponseWriter, r *http.Request) {
		syncHandler(w, r, n)
	})

	handler.HandleFunc(endpointAddPeer, func(w http.ResponseWriter, r *http.Request) {
		addPeerHandler(w, r, n)
	})

	handler.HandleFunc(endpointBlockByNumberOrHash, func(w http.ResponseWriter, r *http.Request) {
		blockByNumberOrHash(w, r, n)
	})

	if isSSLDisabled {
		server := &http.Server{Addr: fmt.Sprintf(":%d", n.info.Port), Handler: handler}

		go func() {
			<-ctx.Done()
			_ = server.Close()
		}()

		err := server.ListenAndServe()
		// This shouldn't be an error!
		if err != http.ErrServerClosed {
			return err
		}

		return nil
	} else {
		certmagic.DefaultACME.Email = sslEmail

		return certmagic.HTTPS([]string{n.info.IP}, handler)
	}
}

func (n *Node) mine(ctx context.Context) error {
	var miningCtx context.Context
	var stopCurrentMining context.CancelFunc

	ticker := time.NewTicker(time.Second * miningIntervalSeconds)

	for {
		select {
		case <-ticker.C:
			go func() {
				if len(n.pendingTXs) > 0 && !n.isMining {
					n.isMining = true

					miningCtx, stopCurrentMining = context.WithCancel(ctx)
					err := n.minePendingTXs(miningCtx)
					if err != nil {
						fmt.Printf("ERROR: %s\n", err)
					}

					n.isMining = false
				}
			}()

		case block, _ := <-n.newSyncedBlocks:
			if n.isMining {
				blockHash, _ := block.Hash()
				fmt.Printf("\nPeer mined next Block '%s' faster :(\n", blockHash.Hex())

				n.removeMinedPendingTXs(block)
				stopCurrentMining()
			}

		case <-ctx.Done():
			ticker.Stop()
			return nil
		}
	}
}

func (n *Node) minePendingTXs(ctx context.Context) error {
	blockToMine := NewPendingBlock(
		n.state.LatestBlockHash(),
		n.state.NextBlockNumber(),
		n.info.Account,
		n.getPendingTXsAsArray(),
	)

	minedBlock, err := Mine(ctx, blockToMine, n.miningDifficulty)
	if err != nil {
		return err
	}

	n.removeMinedPendingTXs(minedBlock)

	err = n.addBlock(minedBlock)
	if err != nil {
		return err
	}

	return nil
}

func (n *Node) removeMinedPendingTXs(block database.Block) {
	if len(block.TXs) > 0 && len(n.pendingTXs) > 0 {
		fmt.Println("Updating in-memory Pending TXs Pool:")
	}

	for _, tx := range block.TXs {
		txHash, _ := tx.Hash()
		if _, exists := n.pendingTXs[txHash.Hex()]; exists {
			fmt.Printf("\t-archiving mined TX: %s\n", txHash.Hex())

			n.archivedTXs[txHash.Hex()] = tx
			delete(n.pendingTXs, txHash.Hex())
		}
	}
}

func (n *Node) ChangeMiningDifficulty(newDifficulty uint) {
	n.miningDifficulty = newDifficulty
	n.state.ChangeMiningDifficulty(newDifficulty)
}

func (n *Node) AddPeer(peer PeerNode) {
	n.knownPeers[peer.TcpAddress()] = peer
}

func (n *Node) RemovePeer(peer PeerNode) {
	delete(n.knownPeers, peer.TcpAddress())
}

func (n *Node) IsKnownPeer(peer PeerNode) bool {
	if peer.IP == n.info.IP && peer.Port == n.info.Port {
		return true
	}

	_, isKnownPeer := n.knownPeers[peer.TcpAddress()]

	return isKnownPeer
}

func (n *Node) AddPendingTX(tx database.SignedTx, fromPeer PeerNode) error {
	txHash, err := tx.Hash()
	if err != nil {
		return err
	}

	txJson, err := json.Marshal(tx)
	if err != nil {
		return err
	}

	err = n.validateTxBeforeAddingToMempool(tx)
	if err != nil {
		return err
	}

	_, isAlreadyPending := n.pendingTXs[txHash.Hex()]
	_, isArchived := n.archivedTXs[txHash.Hex()]

	if !isAlreadyPending && !isArchived {
		fmt.Printf("Added Pending TX %s from Peer %s\n", txJson, fromPeer.TcpAddress())
		n.pendingTXs[txHash.Hex()] = tx
		n.newPendingTXs <- tx
	}

	return nil
}

// addBlock is a wrapper around the n.state.AddBlock() to have a single function for changing the main state
// from the Node perspective, so we can also reset the pending state in the same time.
func (n *Node) addBlock(block database.Block) error {
	_, err := n.state.AddBlock(block)
	if err != nil {
		return err
	}

	// Reset the pending state
	pendingState := n.state.Copy()
	n.pendingState = &pendingState

	return nil
}

// validateTxBeforeAddingToMempool ensures the TX is authentic, with correct nonce, and the sender has sufficient
// funds so we waste PoW resources on TX we can tell in advance are wrong.
func (n *Node) validateTxBeforeAddingToMempool(tx database.SignedTx) error {
	return database.ApplyTx(tx, n.pendingState)
}

func (n *Node) getPendingTXsAsArray() []database.SignedTx {
	txs := make([]database.SignedTx, len(n.pendingTXs))

	i := 0
	for _, tx := range n.pendingTXs {
		txs[i] = tx
		i++
	}

	return txs
}
