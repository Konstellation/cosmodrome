package types

import (
	"github.com/tendermint/tendermint/crypto"

	"github.com/cosmos/cosmos-sdk/x/genaccounts"
)

type NodeConfig struct {
	DirName   string
	DaemonDir string
	CliDir    string
}

type Description struct {
	Moniker  string `json:"moniker"`
	Identity string `json:"identity"`
	Website  string `json:"website"`
	Details  string `json:"details"`
}

type ValidatorKey struct {
	Address      string `json:"address"`
	CoinDelegate int64  `json:"coin_delegate"`
}

type ValidatorInfo struct {
	Name        string       `json:"name"`
	IP          string       `json:"ip"`
	Index       int          `json:"index"`
	Cors        string       `json:"cors"`
	Faucet      bool         `json:"faucet"`
	Key         ValidatorKey `json:"key"`
	Description Description  `json:"description"`
	Config      *Config      `json:"config"`
}

type Validator struct {
	NodeConfig  NodeConfig
	Index       int
	ChainID     string
	Moniker     string
	ID          string
	GenFile     string
	GenAccount  *genaccounts.GenesisAccount
	Memo        string
	Cors        string
	ValPubKey   crypto.PubKey
	IP          string
	Key         ValidatorKey
	Description Description
}
