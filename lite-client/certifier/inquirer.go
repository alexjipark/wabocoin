package certifier

import (
	"github.com/tendermint/tendermint/types"
	"github.com/alexjipark/wabocoin/lite-client"
)

type InquiringCertifier struct {
	Cert 			*DynamicCertifier
	TrustedSeeds	Provider	// These are only properly validated data, from local system
	SeedSourse		Provider	// This is a source of new info, like a node rpc, or other import method
}

func NewInquiring (chainID string, vals *types.ValidatorSet, trusted Provider, source Provider) *InquiringCertifier {
	return &InquiringCertifier{
		Cert: NewDynamic(chainID, vals),
		TrustedSeeds: trusted,
		SeedSourse: source,
	}
}

func (c *InquiringCertifier) Certify(check lite_client.Checkpoint) error {
	err := c.Cert.Certify(check)
	if !IsValidatorChangedErr(err) {
		// 검증인 리스트가 바뀐 것을 제외하고는 전부 Error 처리..
		return err
	}

	err = c.updateToHash(check.Header.ValidatorsHash)

	if err != nil {
		return err
	}
	return c.Cert.Certify(check)
}

// updateToHash gets the validator hash we want to update to
func (c *InquiringCertifier) updateToHash (vhash []byte) error {
	// try to get the match, and update
	seed, err := c.SeedSourse.GetByHash(vhash)

	err = c.Cert.Update(seed.Checkpoint, seed.Validators)

	return err
}

func (c *InquiringCertifier) Update( check lite_client.Checkpoint, vals *types.ValidatorSet) error {
	err := c.Cert.Update(check, vals)
	if err == nil {
		c.TrustedSeeds.StoreSeed(Seed{check, vals})
	}
	return err
}

