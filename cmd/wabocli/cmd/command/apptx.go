package command

import (
	"github.com/tendermint/go-crypto"
	wb "github.com/alexjipark/wabocoin/types"
	"errors"
)

// AppTx Application Transaction structure for client

type AppTx struct {
	chainID string
	signers []crypto.PubKey
	Tx		*wb.AppTx	// [AJ?] why point...
}
// SignBytes returned the unsigned bytes, needing a signature..
func (s *AppTx) SignBytes() []byte {
	return s.Tx.SignBytes(s.chainID)
}

// Sign will add a signature and pubkey
// [AJ?] Depending on the Signable, one may be able to call this multiple times for multisig
// Returns error if called with invalid data or too many times
func (s *AppTx) Sign(pubkey crypto.PubKey, sig crypto.Signature) error {
	if len(s.signers) > 0 {
		return errors.New("AppTx already signed..")
	}

	s.Tx.SetSignature(sig)
	s.signers = []crypto.PubKey{pubkey}

	return nil
}