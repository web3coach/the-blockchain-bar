# The Blockchain Bar

> The source-code for: "Build a Blockchain from Scratch in Go" eBook.

:books: Download the eBook from: [https://web3.coach#book](https://web3.coach#book)

[![book cover](public/img/book_cover2.png)](https://web3.coach#book)

Table of Contents
=================

   * [TBB Training Ledger](#tbb-training-ledger)
      * [Sneak peek to Chapter 13](#sneak-peek-to-chapter-13)
         * [1/4 Check the current blockchain network status](#14-check-the-current-blockchain-network-status)
         * [2/4 Download the pre-compiled blockchain program](#24-download-the-pre-compiled-blockchain-program)
            * [Install](#install)
               * [Download](#download)
                  * [Linux](#linux)
                  * [MacOS](#macos)
               * [Verify the version](#verify-the-version)
         * [3/4 Connect to the training network](#34-connect-to-the-training-network)
         * [4/4 Check the current blockchain network status](#44-check-the-current-blockchain-network-status)
   * [Introduction](#introduction)
      * [How?](#how)
      * [What will you build?](#what-will-you-build)
         * [1) You will build a peer-to-peer system from scratch](#1-you-will-build-a-peer-to-peer-system-from-scratch)
         * [2) You will secure the system with a day-to-day practical cryptography](#2-you-will-secure-the-system-with-a-day-to-day-practical-cryptography)
         * [3) You will implement Bitcoin, Ethereum and XRP backend components](#3-you-will-implement-bitcoin-ethereum-and-xrp-backend-components)
         * [4) You will write unit tests and integration tests for all core components](#4-you-will-write-unit-tests-and-integration-tests-for-all-core-components)
      * [How to use this repository](#how-to-use-this-repository)
      * [Installation](#installation)
      * [Getting started](#getting-started)
   * [Usage](#usage)
      * [Install](#install-1)
      * [CLI](#cli)
         * [Show available commands and flags](#show-available-commands-and-flags)
            * [Show available run settings](#show-available-run-settings)
         * [Run a TBB node connected to the official book's test network](#run-a-tbb-node-connected-to-the-official-books-test-network)
         * [Run a TBB bootstrap node in isolation, on your localhost only](#run-a-tbb-bootstrap-node-in-isolation-on-your-localhost-only)
            * [Run a second TBB node connecting to your first one](#run-a-second-tbb-node-connecting-to-your-first-one)
         * [Create a new account](#create-a-new-account)
      * [HTTP](#http)
         * [List all balances](#list-all-balances)
         * [Send a signed TX](#send-a-signed-tx)
         * [Check node's status (latest block, known peers, pending TXs)](#check-nodes-status-latest-block-known-peers-pending-txs)
      * [Tests](#tests)
   * [Start](#start)
      * [Get the first 7 chapters for FREE](#get-the-first-7-chapters-for-free)
      * [Buy complete eBook](#buy-complete-ebook)
   * [Finish](#finish)
      * [Request 1000 TBB testing tokens](#request-1000-tbb-testing-tokens)

# TBB Training Ledger
## Sneak peek to Chapter 13
You will build all of this from scratch! Follow the next 4 steps and you will be connected to the TBB testing blockchain network in 30 seconds.

### 1/4 Check the current blockchain network status
Go ahead. Try it right now! :cupid:

```bash
curl http://node.tbb.web3.coach/balances/list
```

In case you have the [JQ - CLI JSON processor formatter.](https://github.com/stedolan/jq) installed.

```bash
curl http://node.tbb.web3.coach/balances/list | jq
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
wget "https://github.com/web3coach/the-blockchain-bar/releases/download/1.0.0-alpha/tbb-linux-amd64" -O /usr/local/bin/tbb
```

###### MacOS
```bash
wget "https://github.com/web3coach/the-blockchain-bar/releases/download/1.0.0-alpha/tbb-osx" -O /usr/local/bin/tbb
```

##### Verify the version
```bash
chmod a+x /usr/local/bin/tbb
tbb version

> Version: 1.0.0-beta TBB Training Ledger
```

### 3/4 Connect to the training network from localhost
```bash
tbb run --datadir=$HOME/.tbb --ip=127.0.0.1 --port=8081 --disable-ssl
```

Your blockchain database will synchronize with rest of the students.
```json
Launching TBB node and its HTTP API...
Listening on: 127.0.0.1:8081
Blockchain state:
	- height: 0
	- hash: 0000000000000000000000000000000000000000000000000000000000000000

Searching for new Peers and their Blocks and Peers: 'node.tbb.web3.coach'
Found 1 new blocks from Peer node.tbb.web3.coach
Importing blocks from Peer node.tbb.web3.coach...

Persisting new Block to disk:
	{"hash":"000000a9d18730c133869d175a886d576df5675e0e73900bf072c59047b9d734","block":{"header":{"parent":"0000000000000000000000000000000000000000000000000000000000000000","number":0,"nonce":1925346453,"time":1590684713,"miner":"0x09ee50f2f37fcba1845de6fe5c762e83e65e755c"},"payload":[{"from":"0x09ee50f2f37fcba1845de6fe5c762e83e65e755c","to":"0x22ba1f80452e6220c7cc6ea2d1e3eeddac5f694a","value":5,"nonce":1,"data":"","time":1590684702,"signature":"0JE1yEoA3gwIiTj5ayanUZfo5ZnN7kHIRQPOw8/OZIRYWjbvbMA7vWdPgoqxnhFGiTH7FIbjCQJ25fQlvMvmPwA="}]}}


Searching for new Peers, Blocks: 'node.tbb.web3.coach'
Searching for new Peers, Blocks: 'node.tbb.web3.coach'
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
> c8_transparent_db
> c9_tango
> c10_peer_sync
> c11_consensus
> c12_crypto
> c13_training_network
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

# Usage

## Install
```
go install ./cmd/...
```

## CLI
### Show available commands and flags
```bash
tbb help
```

#### Show available run settings
```bash
tbb run --help

Launches the TBB node and its HTTP API.

Usage:
  tbb run [flags]

Flags:
      --bootstrap-account string   default bootstrap Web3Coach's Genesis account with 1M TBB tokens (default "0x09ee50f2f37fcba1845de6fe5c762e83e65e755c")
      --bootstrap-ip string        default bootstrap Web3Coach's server to interconnect peers (default "node.tbb.web3.coach")
      --bootstrap-port uint        default bootstrap Web3Coach's server port to interconnect peers (default 443)
      --datadir string             Absolute path to your node's data dir where the DB will be/is stored
      --disable-ssl                should the HTTP API SSL certificate be disabled? (default false)
  -h, --help                       help for run
      --ip string                  your node's public IP to communication with other peers (default "127.0.0.1")
      --miner string               your node's miner account to receive the block rewards (default "0x0000000000000000000000000000000000000000")
      --port uint                  your node's public HTTP port for communication with other peers (configurable if SSL is disabled) (default 443)
```

### Run a TBB node connected to the official book's test network 
If you are running the node on your localhost, just disable the SSL with `--disable-ssl` flag.

```
tbb version
> Version: 1.2.0-alpha TBB Training Ledger - HTTPS

tbb run --datadir=$HOME/.tbb --ip=127.0.0.1 --port=8081 --miner=0x_YOUR_WALLET_ACCOUNT --disable-ssl
```

### Run a TBB bootstrap node in isolation, on your localhost only
```
tbb run --datadir=$HOME/.tbb_boostrap --ip=127.0.0.1 --port=8080 --bootstrap-ip=127.0.0.1 --bootstrap-port=8080 --disable-ssl
```

#### Run a second TBB node connecting to your first one
```
tbb run --datadir=$HOME/.tbb --ip=127.0.0.1 --port=8081 --bootstrap-ip=127.0.0.1 --bootstrap-port=8080 --disable-ssl
```

### Create a new account
```
tbb wallet new-account --datadir=$HOME/.tbb 
```

### Run a TBB node in production
The default node's HTTP port is 443. The SSL certificate is generated automatically as long as the DNS A/AAAA records point at your server.

#### Production Server
Example how the official TBB bootstrap node is launched. Customize the `--datadir`, `--miner`, and `--ip` values to match your server.

```bash
/usr/local/bin/tbb run --datadir=/home/ec2-user/.tbb --miner=0x09ee50f2f37fcba1845de6fe5c762e83e65e755c --ip=node.tbb.web3.coach --port=443 --ssl-email=lukas@web3.coach --bootstrap-ip=node.tbb.web3.coach --bootstrap-port=443 --bootstrap-account=0x09ee50f2f37fcba1845de6fe5c762e83e65e755c
```

## HTTP
### List all balances
```
curl http://localhost:8080/balances/list | jq
```

### Send a signed TX
```
curl --location --request POST 'http://localhost:8080/tx/add' \
--header 'Content-Type: application/json' \
--data-raw '{
	"from": "0x22ba1f80452e6220c7cc6ea2d1e3eeddac5f694a",
	"from_pwd": "security123",
	"to": "0x6fdc0d8d15ae6b4ebf45c52fd2aafbcbb19a65c8",
	"value": 100
}'
```

### Check node's status (latest block, known peers, pending TXs)
```
curl http://localhost:8080/node/status | jq
```

## Tests
Run all tests with verbosity but one at a time, without timeout, to avoid ports collisions:
```
go test -v -p=1 -timeout=0 ./...
```

Run an individual test:
```
go test -timeout=0 ./node -test.v -test.run ^TestNode_Mining$
```

**Note:** Majority are integration tests and take time. Expect the test suite to finish in ~30 mins.

# Start
## Get the first 7 chapters for FREE
> Give it a try.

:inbox_tray: Download the newsletter-edition from: [https://web3.coach](https://web3.coach)

## Buy complete eBook
> Time to expand your programming career!

:books: Purchase from Gumroad: [https://gumroad.com/l/build-a-blockchain-from-scratch-in-go](https://gumroad.com/l/build-a-blockchain-from-scratch-in-go)

# Finish
## Request 1000 TBB testing tokens

Write a tweet and let me know how did you like this book! Tag me in it [@Web3Coach](https://twitter.com/Web3Coach) and include your account address 0xYOUR_ADDRESS.

See you on Twitter - [@Web3Coach.](https://twitter.com/Web3Coach)
