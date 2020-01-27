package node

import (
	"context"
	"fmt"
	"github.com/web3coach/the-blockchain-bar/database"
	"github.com/web3coach/the-blockchain-bar/fs"
	"math/rand"
	"time"
)

type PendingBlock struct {
	parent database.Hash
	number uint64
	time   uint64
	txs    []database.Tx
}

func NewPendingBlock(parent database.Hash, number uint64, txs []database.Tx) PendingBlock {
	return PendingBlock{parent, number, uint64(time.Now().Unix()), txs}
}

func Mine(ctx context.Context, pb PendingBlock) (database.Block, error) {
	if len(pb.txs) == 0 {
		return database.Block{}, fmt.Errorf("mining empty blocks is not allowed")
	}

	start := time.Now()
	attempt := 0
	var block database.Block
	var hash database.Hash
	var nonce uint32

	for !database.IsBlockHashValid(hash) {
		select {
		case <-ctx.Done():
			fmt.Println("Mining cancelled!")

			return database.Block{}, fmt.Errorf("mining cancelled. %s", ctx.Err())
		default:
		}

		attempt++
		nonce = generateNonce()

		if attempt%1000000 == 0 || attempt == 1 {
			fmt.Printf("Mining %d Pending TXs. Attempt: %d\n", len(pb.txs), attempt)
		}

		block = database.NewBlock(pb.parent, pb.number, nonce, pb.time, pb.txs)
		blockHash, err := block.Hash()
		if err != nil {
			return database.Block{}, fmt.Errorf("couldn't mine block. %s", err.Error())
		}

		hash = blockHash
	}

	fmt.Printf("\nMined new Block '%x' using PoW🎉🎉🎉%s:\n", hash, fs.Unicode("\\U1F389"))
	fmt.Printf("\tHeight: '%v'\n", pb.number)
	fmt.Printf("\tNonce: '%v'\n", nonce)
	fmt.Printf("\tCreated: '%v'\n", pb.time)
	fmt.Printf("\tParent: '%v'\n\n", pb.parent.Hex())

	fmt.Printf("\tAttempt: '%v'\n", attempt)
	fmt.Printf("\tTime: %s\n\n", time.Since(start))

	return block, nil
}

func generateNonce() uint32 {
	rand.Seed(time.Now().UTC().UnixNano())

	return rand.Uint32()
}
