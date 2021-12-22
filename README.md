# Konstellation network

## Localnet

Run in shell from project dir
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
Run in shell from project dir
#### Create testnet
```shell script
./scripts/testnet.sh create
```
#### Deploy testnet
```shell script
./scripts/testnet.sh deploy

wget https://gist.github.com/Konstellation/b9168ec665bf8991a1cd20fd999452fa/raw/2c53c4c2fa0d90e7a10a6b7f2b5e28c35bec73d2/linux_amd64.tar.gz
tar -xvzf linux_amd64.tar.gz
cp ./linux_amd64/* /usr/local/bin

konstellationlcd rest-server --chain-id darchub --laddr tcp://0.0.0.0:1317


konstellation init kn --chain-id darchub
wget -O ~/.konstellation/config/config.toml https://gist.github.com/Konstellation/b9168ec665bf8991a1cd20fd999452fa/raw/2c53c4c2fa0d90e7a10a6b7f2b5e28c35bec73d2/config.toml
wget https://gist.githubusercontent.com/Konstellation/b9168ec665bf8991a1cd20fd999452fa/raw/2c53c4c2fa0d90e7a10a6b7f2b5e28c35bec73d2/genesis.json
wget https://gist.github.com/Konstellation/b9168ec665bf8991a1cd20fd999452fa/raw/2c53c4c2fa0d90e7a10a6b7f2b5e28c35bec73d2/konstellation.toml

rm -rdf go
rm -rdf ./konstellation
rm -rdf ./konstellationcli
rm /usr/local/bin/konstellation
rm /usr/local/bin/konstellationcli
rm /usr/local/bin/konstellationlcd
```
#### Run testnet nodes on the server side
```shell script
./scripts/testnet.sh run
```
#### Copy config and genesis to konstellation dir
```shell script
./scripts/testnet.sh copy
```

#### Run full node
```shell script
./scripts/fullnode.sh
```
