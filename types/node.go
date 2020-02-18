package types

import (
	"errors"
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

type KeyInfo struct {
	Address     string `json:"address"`
	CoinGenesis int64  `json:"coin_genesis"`
}

type NodeInfo struct {
	Name        string      `json:"name"`
	IP          string      `json:"ip"`
	Index       int         `json:"index"`
	Cors        string      `json:"cors"`
	Faucet      bool        `json:"faucet"`
	Key         Key         `json:"key"`
	Description Description `json:"description"`
}

type Node struct {
	Config      NodeConfig
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
	Key         Key
	Description Description
}

type Key struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Mnemonic string `json:"mnemonic"`
}

type KeyStorage struct {
	Keys map[string]*Key `json:"keys"`
}

func (ks *KeyStorage) GetKey(address string) (*Key, error) {
	key, ex := ks.Keys[address]
	if !ex {
		return nil, errors.New("key does not exist")
	}

	return key, nil
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

type GenAccount struct {
	Address     string `json:"address"`
	CoinGenesis int64  `json:"coin_genesis"`
}

type NetConfig struct {
	Type        string           `json:"type"` // localnet,testnet
	GenAccounts []*GenAccount    `json:"gen_accounts"`
	Validators  []*ValidatorInfo `json:"validators"`
}
