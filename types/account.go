package types

import (
	"github.com/tendermint/go-crypto"
	"fmt"
)

type Account struct {
	Pubkey  		crypto.PubKey	`json:"pubkey"`
	Sequence 		int 			`json:"sequence"`
	Balances 		Coins     		`json:"balance"`
}

func (acc *Account) String() string {	// [AJ?] why pointer?
	if acc == nil {
		return "nil-Account"
	}
	return fmt.Sprintf("%v %v %v", acc.Pubkey, acc.Sequence, acc.Balances)
}

func AccountKey(account []byte ) []byte {
	return append([]byte("wabo/a/"), account...)
}

