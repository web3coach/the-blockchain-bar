# Deployment

## Deploy to the official TBB server
```
ssh tbb
sudo supervisorctl stop tbb
sudo rm /usr/local/bin/tbb
sudo rm /home/ec2-user/tbb
xgo --targets=linux/amd64 ./cmd/tbb
scp -i ~/.ssh/tbb_aws.pem tbb-linux-amd64 ec2-user@ec2-3-127-248-10.eu-central-1.compute.amazonaws.com:/home/ec2-user/tbb
ssh tbb
chmod a+x /home/ec2-user/tbb
sudo ln -s /home/ec2-user/tbb /usr/local/bin/tbb
tbb version
sudo supervisorctl start tbb
```