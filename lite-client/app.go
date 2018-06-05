package lite_client

import (
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/go-data"
	"github.com/pkg/errors"
)

//========== Prover for the lite client =========//

type AppProver struct {
	node client.Client
}

func NewAppProver (node client.Client) AppProver {
	return AppProver{node:node}
}

func (a AppProver) Get(key []byte, height uint64) (Proof, error) {

	resp, err := a.node.ABCIQuery("/key", key)
	if err != nil {
		return nil, err
	}

	if resp.Response.Code == 0 {
		return nil, errors.Errorf("Query error :%d, %s", resp.Response.Code, resp.Response.Log)
	}

	if len(resp.Response.Value) == 0 || len(resp.Response.Proof) == 0 ||
		len(resp.Response.Key) == 0 {
		return nil, ErrNoData()
	}

	if height != 0 && int64(height) != resp.Response.Height {

		return nil, errors.Errorf("Height Wrong")
	}

	proof := AppProof{
		Height: uint64(resp.Response.Height),
		Key: resp.Response.Key,
		Value: resp.Response.Value,
		Proof: resp.Response.Proof,
	}

	return proof, nil
}

func (a AppProver) Unmarshal(data []byte) (Proof, error) {
	return nil, nil
}

//=========== Proof for the lite client ===========//
type AppProof struct {
	Height 	uint64
	Key 	data.Bytes
	Value 	data.Bytes
	Proof 	data.Bytes
}

func (p AppProof) BlockHeight() uint64 {
	return p.Height
}

func (p AppProof) Validate(point Checkpoint) error {
	return nil
}

func (p AppProof) Marshal() ([]byte, error) {
	return nil, nil
}

func (p AppProof) Data() []byte {
	return p.Value
}