package lite_client

import "github.com/tendermint/tendermint/types"

// Basepoint for proving anything on blockchain..
// contains a signed header.
// If the signatures are valid and > 2/3 of the known set of validators
// We can store this checkpoint and use it to prove any number of aspects of the system
// : txs, abci state, validator sets??!!
type Checkpoint struct {
	Header 	*types.Header 	`json:"header"`	// why pointer?
	Commit 	*types.Commit  	`json:"commit"`
}

// ValidateBasic does basic consistency checks and make sure
// the headers and commits are all consistent and refer to our chain..
//
// Make sure to use a Verifier to validate the signatures actually provide
// a significantly strong proof for this header's validity
func (c Checkpoint) ValidateBasic(chainID string) error {
	return nil
}



