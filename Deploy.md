# Deployment

## Deploy to the official TBB server
```
ssh tbb
sudo supervisorctl stop tbb
sudo rm /usr/local/bin/tbb
sudo rm /home/ec2-user/tbb
cd /home/ec2-user/go/src/github.com/web3coach/the-blockchain-bar
git checkout <X>
git pull
make install
tbb version
sudo supervisorctl start tbb
```