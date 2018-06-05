package certifier

import (
	"github.com/tendermint/tendermint/types"
	"github.com/alexjipark/wabocoin/lite-client"
	"fmt"
)

// DynamicCertifier uses a StaticCertifier to evaluate the checkpoint
// but allows for a chage, if we present enough proof

type DynamicCertifier struct {
	Cert 			*StaticCertifier
	LastHeight 		int
}

func NewDynamic(chainID string, vals *types.ValidatorSet) *DynamicCertifier {
	return &DynamicCertifier{
		Cert: NewStatic( chainID, vals),
		LastHeight:0,
	}
}

func (c *DynamicCertifier) Certify (checkpoint lite_client.Checkpoint) error {
	err := c.Cert.Certify(checkpoint)
	if err == nil {
		// update last seen height if input is valid..
		c.LastHeight = int(checkpoint.Header.Height)
	}

	return err
}

// Update will verify if this is a valid change
// and update the certifying validator set if safe to do so.
func (c *DynamicCertifier) Update(check lite_client.Checkpoint, vset *types.ValidatorSet) error {
	// ignore all checkpoints in the past
	if check.Header.Height <= int64(c.LastHeight) {
		return nil	// ErrPastTime
	}

	// first, verify if the input is self-consistent
	err := check.ValidateBasic(c.Cert.ChainID)
	if err != nil {
		return err
	}

	// [AJ?] make sure not too much change...
	// meaning this commit would be proved by the current known validator set
	// as well as the new set.. ??
	err = VerifyCommitAny(c.Cert.VSet, vset, c.Cert.ChainID, check.Commit.BlockID,
							int(check.Header.Height), check.Commit)
	if err != nil {
		return err
	}

	// looks good. we can update
	c.Cert = NewStatic(c.Cert.ChainID, vset)
	c.LastHeight = int(check.Header.Height)

	return nil
}


// VerifyCommitAny will check to see if the wet woud be valid with a different validator set
// old - the validator set we know * over 2/3 of the power in old signed this block
// cur - the validator set that signed this block * only votes from old are sufficient for 2/3 majority in new set as well
// This means that
// 10% of the validator set cannot declare themselves kings
// If validator set is 3x old set, we need more proof to trust
func VerifyCommitAny (old, cur *types.ValidatorSet, chainID string, blockID types.BlockID, height int, commit *types.Commit) error {
	if cur.Size() != len(commit.Precommits) {
		return nil	// error..
	}
	if int64(height) != commit.Height() {
		return nil	// error..
	}

	oldVotingPower := int64(0)
	curVotingPower := int64(0)
	seen := map[int]bool{}
	round := commit.Round()

	for idx, precommit := range commit.Precommits {
		if precommit == nil {

		}
		if precommit.Height != int64(height) {

		}
		if precommit.Round != round {

		}
		if precommit.Type != types.VoteTypePrecommit {

		}
		if !blockID.Equals(precommit.BlockID) {

		}

		//we only grab by address , ignoring unknown validators
		vi, ov := old.GetByAddress(precommit.ValidatorAddress)
		if ov == nil || seen[vi] {
			continue
		}

		precommitSignBytes := []byte(" ") // types.SignBytes(chainID, precommit)
		if !ov.PubKey.VerifyBytes(precommitSignBytes, precommit.Signature ) {
			return fmt.Errorf("Invalid commit - invalid signature")
		}
		//Good Commit
		oldVotingPower += ov.VotingPower

		// check new one
		_, cv := cur.GetByIndex(idx)
		if cv.PubKey.Equals(ov.PubKey) {
			// [AJ?] make sure this is properly set in the current block as well
			curVotingPower += cv.VotingPower
		}
	}

	if oldVotingPower <= old.TotalVotingPower() * 2/3 {
		return fmt.Errorf("insufficient voting power from old validators set")
	}
	if curVotingPower <= cur.TotalVotingPower() * 2/3 {
		return fmt.Errorf("insufficient voting power from current validator set")
	}

	return nil
}

