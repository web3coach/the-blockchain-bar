# The Blockchain Bar eBook

> The source-code for the first 7 chapters of: "Build a Blockchain from Scratch in Go" eBook.

Download the eBook from: [https://web3.coach#book](https://web3.coach#book)

![book cover](public/img/book_cover2.png)

# TBB Training Ledger
## Sneak peek to Chapter 13
You will build all of this from scratch! Follow the next 4 steps and you will be connected to the TBB testing blockchain network in 30 seconds.

### 1/4 Check the current blockchain network status
Go ahead. Try it right now! :cupid:

```bash
curl -X GET http://node.tbb.web3.coach:8080/balances/list
```

In case you have the [JQ - CLI JSON processor formatter.](https://github.com/stedolan/jq) installed.

```bash
curl -X GET http://node.tbb.web3.coach:8080/balances/list | jq
```

```json
{
  "block_hash": "000000a9d18730c133869d175a886d576df5675e0e73900bf072c59047b9d734",
  "balances": {
    "0x09ee50f2f37fcba1845de6fe5c762e83e65e755c": 1000095,
    "0x22ba1f80452e6220c7cc6ea2d1e3eeddac5f694a": 5
  }
}
```

### 2/4 Download the pre-compiled blockchain program
Together with all students, we will use this custom build blockchain for educational purposes.

#### Install
##### Download
###### Linux
```bash
wget "https://github.com/web3coach/the-blockchain-bar-newsletter-edition/releases/download/1.0.0-alpha/tbb-linux-amd64" -O /usr/local/bin/tbb
```

###### MacOS
```bash
wget "https://github.com/web3coach/the-blockchain-bar-newsletter-edition/releases/download/1.0.0-alpha/tbb-osx" -O /usr/local/bin/tbb
```

##### Verify the version
```bash
chmod a+x /usr/local/bin/tbb
tbb version

> Version: 1.0.0-beta TBB Training Ledger
```

### 3/4 Connect to the training network
```bash
tbb run --datadir=$HOME/.tbb --ip=127.0.0.1 --port=8081
```

Your blockchain database will synchronize with rest of the students.
```json
Launching TBB node and its HTTP API...
Listening on: 127.0.0.1:8081
Blockchain state:
	- height: 0
	- hash: 0000000000000000000000000000000000000000000000000000000000000000

Searching for new Peers and their Blocks and Peers: 'node.tbb.web3.coach:8080'
Found 1 new blocks from Peer node.tbb.web3.coach:8080
Importing blocks from Peer node.tbb.web3.coach:8080...

Persisting new Block to disk:
	{"hash":"000000a9d18730c133869d175a886d576df5675e0e73900bf072c59047b9d734","block":{"header":{"parent":"0000000000000000000000000000000000000000000000000000000000000000","number":0,"nonce":1925346453,"time":1590684713,"miner":"0x09ee50f2f37fcba1845de6fe5c762e83e65e755c"},"payload":[{"from":"0x09ee50f2f37fcba1845de6fe5c762e83e65e755c","to":"0x22ba1f80452e6220c7cc6ea2d1e3eeddac5f694a","value":5,"nonce":1,"data":"","time":1590684702,"signature":"0JE1yEoA3gwIiTj5ayanUZfo5ZnN7kHIRQPOw8/OZIRYWjbvbMA7vWdPgoqxnhFGiTH7FIbjCQJ25fQlvMvmPwA="}]}}


Searching for new Peers, Blocks: 'node.tbb.web3.coach:8080'
Searching for new Peers, Blocks: 'node.tbb.web3.coach:8080'
```

### 4/4 Check the current blockchain network status
This time from your OWN, fully synchronized blockchain node running directly on your computer.

```bash
curl -X GET http://127.0.0.1:8081/balances/list | jq
```

```json
{
  "block_hash": "000000a9d18730c133869d175a886d576df5675e0e73900bf072c59047b9d734",
  "balances": {
    "0x09ee50f2f37fcba1845de6fe5c762e83e65e755c": 1000095,
    "0x22ba1f80452e6220c7cc6ea2d1e3eeddac5f694a": 5
  }
}
```

:star: All of this and much more, you will build from scratch!

# Introduction

Hi :wave:,

With Web 3.0 and blockchain becoming more mainstream every day, do you know what blockchain is? Do you know its technical advantages and use-cases?

**The goal of this tutorial is to introduce blockchain technology from a technical perspective by building one from scratch.**

Forget everything you've heard about blockchain from social media. Now, you will build a blockchain system from ground zero to really understand the ins and outs of this peer-to-peer, distributed technology.

Afterwards, make your own mind up about its future, advantages and shortcomings. 

Spoiler alert: you will fall in love with programming blockchain software. :smiling_imp:

## How?

You will follow the story of a software developer who is looking to revolutionize his local bar by implementing blockchain technology for its payment system.

Although blockchain has several undeniable use-cases, at the moment, the number one application is payments. This is because banks are still running on an inefficient, 40 year old infrastructure powered by CSV files and FTP.

The story comes with a lot of fun and intriguing facts about the overall blockchain ecosystem and different protocols such as Bitcoin, Ethereum and XRP.

## What will you build?

Chapter by chapter, you will build a full peer-to-peer, autonomous blockchain system in Go and **learn all standard blockchain components!**

### 1) You will build a peer-to-peer system from scratch

You start with 0 lines of code and end-up with 13+ branches with complete executable source-code.

PS: Don't worry if anything on the screen makes sense yet, it will once you go chapter by chapter; release by release.

![peer-to-peer blockchain system in action](public/img/andrej_babayaga_caesar_sync_p2p.png)

### 2) You will secure the system with a day-to-day practical cryptography

No boring theory. Only modern practices.

![elliptic curve cryptography](public/img/andrej_babayaga_crypto_sign_summary.png)

### 3) You will implement Bitcoin, Ethereum and XRP backend components 

From diagrams of mining algorithms to actual, implemented and working crypto wallets for storing the mined tokens and all other fundamental components that make blockchain special.

![decentralized consensus](public/img/mining_p2p.png)

### 4) You will write unit tests and integration tests for all core components

You will test your cryptographic functions, a Bitcoin's like Proof of Work mining algorithm and other key components.

![ethereum signature](public/img/test_ethereum_signature.png)

## How to use this repository
Every eBook chapter has a dedicated branch where you can experiment with the code first-hand.

```git
git pull --all
git branch

> c1_genesis_json
> c2_db_changes_txt
> c3_state_blockchain_component
> c4_caesar_transfer
> c5_broken_trust
> c6_immutable_hash
> c7_blockchain_programming_model
```

## Installation

[Open instructions.](./Installation.md)

## Getting started
1. Download the eBook from: [https://web3.coach#book](https://web3.coach#book)
1. Open the book at Chapter 1
1. Checkout the first chapter's branch `c1_genesis_json`

```git
git pull --all

git checkout c1_genesis_json
```

## TBB program usage (first 7 chapters only)
### CLI
Interacting with TBB blockchain using CLI.

### Compile the code
```bash
go install ./cmd/...
```

### Show current program's version
```bash
tbb version
```

### Show blockchain balances of all bar's customers
```bash
tbb balances list
```

### Store a new TX in the DB
```bash
tbb tx add --from=andrej --to=babayaga --value=1000
```

### Store a new Reward TX in the DB
```bash
tbb tx add --from=andrej --to=andrej --value=100 --data=reward
```

## Getting unstuck
Can't understand why is something done in a particular way or crack your way around a specific chapter's topic?

Blockchain is a challenging technology.
   
Write me a DM on Twitter or create a Github Issue, and I will help you move forward on your new blockchain journey!
   
[https://twitter.com/Web3Coach](https://twitter.com/Web3Coach)
