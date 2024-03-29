# Konstellation network

## Localnet

Build konstellation docker image
```shell script
# Clone Konstellation from the latest release found here: https://github.com/konstellation/konstellation/releases
git clone -b <latest_release> https://github.com/konstellation/konstellation
# Enter the folder Konstellation was cloned into
cd konstellation
# Change git release branch to `master`
git checkout master
# build docker image
docker build -t knstld:latest .
```

Run in shell from cosmodrome project dir
#### Create localnet
```shell script
./scripts/localnet.sh create
```
#### Run localnet
```shell script
./scripts/localnet.sh run
```
#### Copy config and genesis to konstellation dir
```shell script
./scripts/localnet.sh copy
```

#### Stop and remove localnet
```shell script
./scripts/localnet.sh stop
./scripts/localnet.sh rm
```

#### With docker-compose
```shell script
./scripts/localnet.sh create

./scripts/localnet.sh copy

Update moniker in `~/.knstld/config/config.toml`

docker-compose up
```

## Testnet
See https://github.com/Konstellation/cosmodrome/blob/master/config/testnet.md
