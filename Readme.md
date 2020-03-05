# The Blockchain Bar

## Install
```
dep ensure
go install ./cmd/...
```

## Usage
### List all possible commands, arguments and configurations
```
tbb help
```

### Run TBB blockchain connected to the official book's test network
```
tbb run --datadir=~/.tbb
```

### Run TBB blockchain in isolation, on your localhost only
```
tbb run --datadir=~/.tbb --bootstrap=""
```

### Create a new account
```
tbb wallet new-account --datadir=~/.tbb 
```

## HTTP Usage
### List all balances
```
curl -X GET http://localhost:8080/balances/list -H 'Content-Type: application/json'
```

### Send and sign a new TX
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

## Compile
To local OS:
```
go install ./cmd/...
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

## Deploy to official TBB server
```
ssh tbb
sudo supervisorctl stop tbb
sudo rm /usr/local/bin/tbb
sudo rm /home/ec2-user/tbb
xgo --targets=linux/amd64 ./cmd/tbb
scp -i ~/.ssh/tbb_aws.pem tbb-linux-amd64 ec2-user@ec2-18-184-213-146.eu-central-1.compute.amazonaws.com:/home/ec2-user/tbb
ssh tbb
chmod a+x /home/ec2-user/tbb
sudo ln -s /home/ec2-user/tbb /usr/local/bin/tbb
tbb version
sudo supervisorctl start tbb
```
