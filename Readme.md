# The Blockchain Bar
> Build a custom blockchain in Go from scratch.

The source-code for the first 6 chapters of: "The Blockchain Way of Programming".

Download the eBook from: https://web3.coach

![book cover](./book_cover.png)

## How to use this repository
Every eBook chapter has a dedicated branch where you can experiment with the code first-hand.

```git
git branch

> c1_genesis_json
> c2_db_changes_txt
> c3_state_blockchain_component
> c4_caesar_transfer
> c5_broken_trust
> c6_immutable_hash
```

## Installation

### Install Go 1.13 or higher
Follow the official docs or use your favorite dependency manager
to install Go: [https://golang.org/doc/install](https://golang.org/doc/install)

Verify your `$GOPATH` is correctly set before continuing.

### Setup this repository

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

#### Using Go get
```bash
go get -u github.com/web3coach/the-blockchain-way-of-programming-newsletter-edition
```

## Getting started
1. Open the eBook at Chapter 1
1. Checkout the first chapter's branch

```git
git pull --all

git checkout c1_genesis_json
```

Read, experiment with the code and, most importantly, have fun!

## Getting unstuck
Can't understand why is something done in a particular way or crack your way around a specific chapter's topic?

Blockchain is a challenging technology.
   
As I promised, you have my full support. You are not alone in this. Write me a DM on LinkedIn, and I will help you figure it out and move forward on your new journey :)
   
[https://www.linkedin.com/in/llukac](https://www.linkedin.com/in/llukac)
