# The Blockchain Bar
> A custom build blockchain in Go.

The source-code for the first 6 chapters of: "The Blockchain Way of Programming".

Download the eBook from: https://web3.coach

## Foreword
Welcome to the blockchain world!

You made the right choice of learning blockchain development and expanding your programming career.

**This repository is just an extract from the private Github repository and it contains roughly 5-10% of the overall full source code that you will get access to after purchasing the complete eBook once it's ready in a few weeks time.**

Enjoy your blockchain programming journey! 

## How to use this repository
Every eBook chapter has a dedicated branch where you can experiment with the code first-hand.

```git
git branch

> c1_genesis_json
> c2_db_changes_txt
> c3_state_blockchain_component
> c4...
> c5...
```

## Installation

### Install Go
Follow the official docs or use your favorite dependency manager
to install Go: [https://golang.org/doc/install](https://golang.org/doc/install)

Verify your `$GOPATH` is correctly set before continuing.

### Setup this Repository

Go is bit picky about where you store your repositories.

The convention is to store:
- the source code inside the `$GOPATH/src`
- the compiled program binaries inside the `$GOPATH/bin`

You can `clone` the repository or use `go get` to install it.

#### Using Git
```bash
mkdir -p $GOPATH/src/github.com/web3coach
cd $GOPATH/src/github.com/web3coach

git clone git@github.com:web3coach/the-blockchain-way-of-programming-newsletter-edition.git
```

PS: Make sure you actually clone it inside the `src/github.com/web3coach` directory, not your own, otherwise it won't compile. Go rules.

#### Using Go Get
```bash
go get -u github.com/web3coach/the-blockchain-way-of-programming-newsletter-edition
```

## Getting Started
1. Open the eBook at Chapter 1
1. Checkout the first chapter's branch

```git
git checkout c1_genesis_json
```

Read, experiment with the code and, most importantly, have fun!

## CLI
Interacting with blockchain using CLI.

### Show current program's version
```bash
tbb help
```

### Show blockchain balances of all bar's customers
```bash
tbb balances
```

### Store a new TX in the DB
```bash
tbb tx add --from=andrej --to=babayaga --value=1000
```

### Store a new Reward TX in the DB
```bash
tbb tx add --from=andrej --to=andrej --value=100 --data=reward
```

## Getting Unstuck
Can't understand why is something done in a particular way or crack your way around a specific chapter's topic?

Blockchain is a challenging technology.
   
As I promised, you have my full support. You are not alone in this. Write me a DM on LinkedIn, and I will help you figure it out and move forward on your new journey :)
   
[https://www.linkedin.com/in/llukac](https://www.linkedin.com/in/llukac)