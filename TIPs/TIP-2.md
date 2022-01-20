# The Blockchain Bar UI - Database Explorer, Faucet, Wallet
## Vision
To help developers practice blockchain programming in a test-network by creating new features, API endpoints and database indexing.

To help developers learn what's happening behind the scenes in the blockchain network by providing several graphical components painting a better picture on:
- how the blocks are chained
- how account balances are calculated and validated
- what are the transaction attributes
- how to create a transaction
- how to sign a transaction
- how a transaction waits in a Mempool until is mined
- how the mining process works

// TODO: Add more of your ideas here.

**PS for frontend devs: We will need a new `web` directory in this repository with all the frontend code for Explorer, Faucet and Wallet.**

## Database Explorer
### Idea - Block Traversal
#### User Story
As a User I want to traverse database blocks,
so I can see all The Blockchain Bar's users activity and what happened when.

#### Backend Proposal
Add a new API endpoint `/block/$height` to `./node/http_routes.go` to retrieve block by number.

Add a new API endpoint `/block/$hash` to `./node/http_routes.go`  to retrieve block by hash.

Both endpoints will return a Block as JSON:
```
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
```

The easiest way to achieve this ATM will be by modifying the `func GetBlocksAfter(blockHash Hash, dataDir string) ([]Block, error) {` and adding a new argument `limit uint` to the function. As well as creating another getter to the `database` package like `func GetBlocksByNumberAfter(blockHash Hash, limit uint, dataDir string) ([]Block, error) {` and abstracting the common code shared between these two functions (the already implement blocks retrieval from disk) to a third function.

#### Frontend Proposal
// TODO: Add here some UI ideas how this could be displayed in the UI.

---

### Idea - Query Balance
#### User Story
As a User I want to check a balance of an account.

#### Backend Proposal

#### Frontend Proposal

---

### Idea - Visualize Mempool
#### User Story
As a User I want to see what's the difference between a database and a Mempool.

#### Backend Proposal

#### Frontend Proposal

---

### Idea - Visualize P2P with Animations
#### User Story
As a User I want to see how transactions are broadcasted over p2p,
so I have a better idea what's happening after a transaction is broadcasted.

#### Backend Proposal

#### Frontend Proposal

---

## Wallet
### Idea - Create Account
#### User Story
As a User I want to create my The Blockchain Bar account,
so I can receive and send testing tokens.

#### Backend Proposal

#### Frontend Proposal

---

### Idea - Create a Transaction and Sign it on Frontend
#### User Story
As a User I want to learn how to construct a valid Transactions,
so I learn about gas fees, nonce and crypto signing.

#### Backend Proposal

#### Frontend Proposal

---

## Faucet
### Idea - Faucet for Giving Away Free Tokens
#### User Story
As a User I want to request free testing tokens,
so I can experiment around with my Wallet.

#### Backend Proposal

#### Frontend Proposal

--- 

### Idea - Visualize Mempool
#### User Story
As a User I want to see what's the difference between a database and a Mempool.

#### Backend Proposal

#### Frontend Proposal

--- 

## Currently Available API endpoints:
The below list is very limited and would need a lot of new API endpoints to support all above UI features. Great opportunity to practice Go programming. 

### Query the latest balances of all accounts
https://node.tbb.web3.coach/balances/list

```json
{
  "block_hash": "00000050159ee00da041cb8bd1a1974ea80da53f8489ade6c4886dcf48205ca9",
  "balances": {
    "0x09ee50f2f37fcba1845de6fe5c762e83e65e755c": 988845,
    "0x0fc524c45b51e3215701ef7c12e58c781b0642ac": 1000,
    ...
    ...
    "0x9f23369f02924f94765711bb09ebe1937123e37d": 1000,
    "0xdd1fbd71d7f78a810ae10829220fbb7bd2f1818b": 150,
    "0xf04679700642fd1455782e1acc37c314aab1a847": 1550,
    "0xf336ed6a7c7529df70552377d641088be4dc4519": 1000
  }
}
```

### Query the latest node status
https://node.tbb.web3.coach/node/status

```json
{
  "block_hash": "00000050159ee00da041cb8bd1a1974ea80da53f8489ade6c4886dcf48205ca9",
  "block_number": 36,
  "peers_known": {
    "node.tbb.web3.coach:443": {
      "ip": "node.tbb.web3.coach",
      "port": 443,
      "is_bootstrap": true,
      "account": "0x09ee50f2f37fcba1845de6fe5c762e83e65e755c",
      "node_version": ""
    }
  },
  "pending_txs": [],
  "node_version": "1.9.0-alpha 959541 TX Gas",
  "account": "0x09ee50f2f37fcba1845de6fe5c762e83e65e755c"
}
```

### Unlock wallet to sign a transaction and broadcast it
```
curl --location --request POST 'http://localhost:8080/tx/add' \
--header 'Content-Type: application/json' \
--data-raw '{
	"from": "0x22ba1f...",
	"from_pwd": "",
	"to": "0x6fdc0d...",
	"value": 100,
	"gas": 21,
	"gas_price": 1
}'
```