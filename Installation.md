## Installation

### Install Go 1.13 or higher
Follow the official docs or use your favorite dependency manager
to install Go: [https://golang.org/doc/install](https://golang.org/doc/install)

Verify your `$GOPATH` is correctly set before continuing!

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

git clone git@github.com:web3coach/the-blockchain-bar.git
```

PS: Make sure you actually clone it inside the `src/github.com/web3coach` directory, not your own, otherwise it won't compile. Go rules.

#### Using Go get
```bash
go get -u github.com/web3coach/the-blockchain-bar
```