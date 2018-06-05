package lite_client

// Proof is a generic interface for data along with the cryptographic proof
// of data's validity, tied to a checkpoint..
type Proof interface {
	BlockHeight() uint64
	Validate(Checkpoint) error	// Checkpoint is validated and proper height
	Marshal() ([]byte, error) 	// Marshal prepares for storage
	Data() []byte 				// extract the query result we want to see..
}

// Prover is anything that can provide proofs..
// Such as a AppProver (for merkle proof of app state)
// or TxProver (for merkle proof that a tx is in a block)

type Prover interface {
	// Get returns the key for the given block height
	// [AJ?] The prover should accept h=0 for latest height
	Get(key []byte, height uint64) (Proof, error)
	Unmarshal([]byte) (Proof, error)
}