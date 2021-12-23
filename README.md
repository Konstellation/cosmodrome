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
See https://github.com/Konstellation/cosmodrome/blob/master/config/testnet.md
