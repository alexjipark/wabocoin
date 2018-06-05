package certifier

import (
	"github.com/tendermint/tendermint/types"
	"github.com/alexjipark/wabocoin/lite-client"
	"bytes"
	"github.com/pkg/errors"
)

// StaticCertifier assumes a static set of validators, set on initialization and check against them

type StaticCertifier struct {
	ChainID			string
	VSet 			*types.ValidatorSet
	vHash 			[]byte
}

func NewStatic(chainID string, vals *types.ValidatorSet) *StaticCertifier {
	return &StaticCertifier{
		ChainID: chainID,
		VSet: vals,
	}
}

func (c *StaticCertifier) Hash() []byte {
	if len(c.vHash) == 0 {
		c.vHash = c.VSet.Hash()
	}

	return c.vHash
}

func (c *StaticCertifier) Certify(check lite_client.Checkpoint) error {
	// Do basic sanity check..
	err := check.ValidateBasic(c.ChainID)
	if err != nil {
		return nil
	}

	// Make sure it has the same validator set as we have
	if !bytes.Equal(c.Hash(), check.Header.ValidatorsHash) {
		return errors.Errorf("not equal validator set..")
	}

	// then, make sure we have the proper signatures for this..
	err = c.VSet.VerifyCommit(c.ChainID, check.Commit.BlockID,
		check.Header.Height, check.Commit)

	return errors.WithStack(err)

}