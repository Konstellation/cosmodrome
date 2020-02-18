package types

import (
	"errors"
)

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
