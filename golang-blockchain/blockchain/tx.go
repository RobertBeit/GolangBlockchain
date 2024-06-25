package blockchain

import (
	"bytes"

	"github.com/tensor-programming/golang-blockchain/wallet"
)

type TxOutput struct {
	Value      int
	PubKeyHash string
}

type TxInput struct {
	ID        []byte
	Out       int
	Signature []byte
	PubKey    []byte
}

func (in *TxInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := wallet.PublicKeyHash(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}
