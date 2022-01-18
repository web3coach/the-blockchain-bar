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

## Database Explorer
### Ideas
As a User I want to traverse database blocks,
so I can see all The Blockchain Bar's users activity and what happened when.

As a User I want to check a balance of an account.

As a User I want to see what's the difference between a database and a Mempool.

As a User I want to see how transactions are broadcasted over p2p,
so I have a better idea what's happening after a transaction is broadcasted.

### Frontend + Backend Proposal
// TODO: Add here your proposal how the UI could work and what backend API endpoints you need.

## Wallet
### Ideas
As a User I want to create my The Blockchain Bar account,
so I can receive and send testing tokens.

As a User I want to learn how to construct a valid Transactions,
so I learn about gas fees, nonce and crypto signing.

### Frontend + Backend Proposal
// TODO: Add here your proposal how the UI could work and what backend API endpoints you need.

## Faucet
### Ideas
As a User I want to request free testing tokens,
so I can experiment around with my Wallet.

### Frontend + Backend Proposal
// TODO: Add here your proposal how the UI could work and what backend API endpoints you need.

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