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

// [AJ] Why this function is needed..
func (s *SendTx) AddSigner(pk crypto.PubKey) {
	var addr []byte
	if len(pk.Bytes()) != 0 {
		addr = pk.Address()
	}

	// set the send address, pubkey if needed..
	in := s.Tx.Inputs
	in[0].Address = addr
	if in[0].Sequence == 1 {
		in[0].PubKey = pk
	}
}

func (s *SendTx) ValidateBasic() error {
	return nil
}