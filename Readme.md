# The Blockchain Bar

## Install
```
go install ./cmd/...
```

## Usage
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
      --bootstrap-port uint        default bootstrap Web3Coach's server port to interconnect peers (default 8080)
      --datadir string             Absolute path to your node's data dir where the DB will be/is stored
  -h, --help                       help for run
      --ip string                  your node's public IP to communication with other peers (default "127.0.0.1")
      --miner string               your node's miner account to receive the block rewards (default "0x0000000000000000000000000000000000000000")
      --port uint                  your node's public HTTP port for communication with other peers (default 8080)
```

### Run a TBB node connected to the official book's test network 
```
tbb version
> Version: 1.0.0-beta TBB Training Ledger

tbb run --datadir=$HOME/.tbb --ip=127.0.0.1 --port=8081 --miner=0x_YOUR_WALLET_ACCOUNT
```

### Run a TBB bootstrap node in isolation, on your localhost only
```
tbb run --datadir=$HOME/.tbb_boostrap --ip=127.0.0.1 --port=8080 --bootstrap-ip=127.0.0.1 --bootstrap-port=8080
```

#### Run a second TBB node connecting to your first one
```
tbb run --datadir=$HOME/.tbb --ip=127.0.0.1 --port=8081 --bootstrap-ip=127.0.0.1 --bootstrap-port=8080
```

### Create a new account
```
tbb wallet new-account --datadir=$HOME/.tbb 
```

## HTTP Usage
### List all balances
```
curl -X GET http://localhost:8080/balances/list -H 'Content-Type: application/json'
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
curl --request GET 'http://localhost:8080/node/status'
```

## Compile
To local OS:
```
go install ./cmd/...
```

To local OS - stripping debugging stuff
```
go install -ldflags="-s -w" ./cmd/...
```

To cross-compile:
```
xgo --targets=linux/amd64 ./cmd/tbb
```

## Tests
Run all tests with verbosity but one at a time, without timeout, to avoid ports collisions:
```
go test -v -p=1 -timeout=0 ./...
```

**Note:** Majority are integration tests and take time. Expect the test suite to finish in ~30 mins.

## How I deploy to my TBB node server
```
ssh tbb
sudo supervisorctl stop tbb
sudo rm /usr/local/bin/tbb
sudo rm /home/ec2-user/tbb
scp -i ~/.ssh/tbb_aws.pem $GOPATH/bin/tbb ec2-user@ec2-3-127-248-10.eu-central-1.compute.amazonaws.com:/home/ec2-user/tbb
ssh tbb
chmod a+x /home/ec2-user/tbb
sudo ln -s /home/ec2-user/tbb /usr/local/bin/tbb
tbb version
sudo supervisorctl start tbb
```
