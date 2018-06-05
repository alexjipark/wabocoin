package command

import (
	"github.com/tendermint/go-crypto"
	"github.com/alexjipark/wabocoin/types"
)

type SendTx struct {
	chainID		string
	signers 	[]crypto.PubKey
	Tx			*types.SendTx
}


