#!/usr/bin/env bash

COMMAND=$1
CHAIN_ID=$2

function usage() {
  echo "Usage:"
  echo "  ./testnet.sh command [chain-id]"
  echo ""
  echo "Command:"
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
    if [[ -f "./scripts/tmp_deploy.sh" ]]; then
      rm ./scripts/tmp_deploy.sh
    fi
    jq -r '
    .validators[] |
    "echo \(.ip);
     ssh -i ~/Documents/.ssh/id_rsa.pub root@\(.ip) rm -rdf /root/.konstellation;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@\(.ip) rm -rdf /root/.konstellationcli;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@\(.ip) rm -rdf /root/.konstellationlcd;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@\(.ip) \"ps ax | grep konstellation | awk '\''{print \\$1}'\'' | xargs kill\";
     scp -i ~/Documents/.ssh/id_rsa.pub -r ./testnet/\(.name)/.konstellation root@\(.ip):/root;
     scp -i ~/Documents/.ssh/id_rsa.pub -r ./testnet/\(.name)/.konstellationcli root@\(.ip):/root;
     ssh -i ~/Documents/.ssh/id_rsa.pub root@\(.ip) \"screen -dmS kn konstellation start\";
     echo "
    ' ./config/testnet.json >> ./scripts/tmp_deploy.sh
    chmod +x ./scripts/tmp_deploy.sh
    ./scripts/tmp_deploy.sh
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
