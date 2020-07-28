#!/usr/bin/env bash

COMMAND=$1
CHAIN_ID=$2

function usage() {
  echo "Usage:"
  echo "  ./testnet.sh command [chain-id]"
  echo ""
  echo "Command:"
  echo "  create    Creates new net conf files "
  echo "  run       Run testnet full node on the server side. "
  echo "  deploy    Deploy testnet to testnodes. "
  echo "  copy      Copy config and genesis to yout konstellation dir. "
  echo ""
}

function create() {
  if [[ -d "testnet" ]]; then
    sudo rm -rdf testnet
  fi

  cosmodrome gn --chain-id "$CHAIN_ID" -n ./config/testnet.json -o ./testnet
}

function deploy() {
  if [[ -f "./config/testnet.json" ]]; then
    if [[ -f "./scripts/deploy_testnet_tmp.sh" ]]; then
      rm ./scripts/deploy_testnet_tmp.sh
    fi
    jq -r '
    .validators[] |
    "echo \(.ip);
     ssh -i ~/.ssh/bldg root@\(.ip) \"ps ax | grep konstellation | awk '\''{print \\$1}'\'' | xargs kill\";
     ssh -i ~/.ssh/bldg root@\(.ip) \"konstellation unsafe-reset-all\";
     ssh -i ~/.ssh/bldg root@\(.ip) \"rm -rdf /root/.konstellation\";
     ssh -i ~/.ssh/bldg root@\(.ip) \"rm -rdf /root/.konstellationcli\";
     ssh -i ~/.ssh/bldg root@\(.ip) \"rm -rdf /root/.konstellationlcd\";
     scp -i ~/.ssh/bldg -r ./testnet/\(.name)/.konstellation root@\(.ip):/root;
     scp -i ~/.ssh/bldg -r ./testnet/\(.name)/.konstellationcli root@\(.ip):/root;
     ssh -i ~/.ssh/bldg root@\(.ip) \"screen -dmS kn konstellation start\";
     ssh -i ~/.ssh/bldg root@\(.ip) \"screen -dmS klcd konstellationlcd rest-server --chain-id darchub --laddr tcp:\/\/0.0.0.0:1317\";
     echo "
    ' ./config/testnet.json >> ./scripts/deploy_testnet_tmp.sh
    chmod +x ./scripts/deploy_testnet_tmp.sh
#    ./scripts/deploy_testnet_tmp.sh
  fi
}

function copy() {
  if [ ! -d $HOME/.konstellation ]; then
    echo "Konstellation config dir does not exist"
    echo "Run konstellation init and then run this script again"
    exit
  fi

  cp -r ./testnet/config/* $HOME/.konstellation/config/
}

if [[ -z ${COMMAND} ]]; then
  error "Command must be set !"
  usage
fi

if [[ -z ${CHAIN_ID} ]]; then
  CHAIN_ID="darchub"
fi

if [[ ! -f "./config/testnet.json" ]]; then
  echo "Nodes config DOSE NOT exist !"
  echo "" >&2
  exit 1
fi

case "${COMMAND}" in
"create")
  create
  ;;
"deploy")
  deploy
  ;;
"copy")
  copy
  ;;
*)
  usage
  echo "" >&2
  exit 1
  ;;
esac
