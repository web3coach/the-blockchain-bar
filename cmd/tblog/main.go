package main

import (
	"github.com/davecgh/go-spew/spew"
)

type Tx struct {
	from  string
	to    string
	value uint64
}

func main() {
	// Extra touch pre-defining the Map length to avoid reallocation
	txMap := make(map[string]Tx, 3)
	txMap["tx1"] = Tx{"andrej", "babayaga", 10}
	
	getTXsAsSlice(txMap)
}

func getTXsAsSlice(txMap map[string]Tx) []Tx {
	// Defines the Slice capacity to match the Map elements count
	txs := make([]Tx, 0, 3)
	spew.Dump(txs[0])
	
	for _, tx := range txMap {
		txs = append(txs, tx)
	}
	
	spew.Dump(txs)
	
	return txs
}