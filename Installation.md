## Installation

### Install Go 1.14 or higher
Follow the official docs or use your favorite dependency manager
to install Go: [https://golang.org/doc/install](https://golang.org/doc/install)

Verify your `$GOPATH` is correctly set before continuing!

### Setup this repository

Go is bit picky about where you store your repositories.

The convention is to store:
- the source code inside the `$GOPATH/src`
- the compiled program binaries inside the `$GOPATH/bin`

#### Using Git
```bash
mkdir -p $GOPATH/src/github.com/web3coach
cd $GOPATH/src/github.com/web3coach

git clone https://github.com/web3coach/the-blockchain-bar.git
```

PS: Make sure you actually clone it inside the `src/github.com/web3coach` directory, not your own, otherwise it won't compile. Go rules.

### Apple Silicon

This project currently depends on [gopsutil](https://github.com/shirou/gopsutil) which is a library that provides a cross-platform interface for querying operating system information. Unfortunately, [Apple Silicon/M1](https://en.wikipedia.org/wiki/Apple_silicon) machines are not yet supported by this library which will result in build failures when trying to compile this project locally. If you are using an Apple M1 machine, it is recommended to follow [the guide](./Docker.md) for using [Docker](https://www.docker.com) for local development.

