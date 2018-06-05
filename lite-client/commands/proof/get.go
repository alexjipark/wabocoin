package proof

import (
	"github.com/alexjipark/wabocoin/lite-client"
	"github.com/spf13/viper"
	"github.com/alexjipark/wabocoin/lite-client/commands"
	"github.com/tendermint/tendermint/rpc/client"
	"fmt"
	"github.com/tendermint/go-wire/data"
	"github.com/tendermint/go-wire"
)

// try to get proof for the given key..
// If successful, it will return the proof and also unserialize Proof.Data into data argument

func GetAndParseAppProof (key []byte, data interface{} ) (lite_client.Proof, error){

	height := GetHeight()
	node := commands.GetNode()	// rpc client..
	prover := lite_client.NewAppProver(node)

	proof, err := GetProof( node, prover, key, height)
	if err != nil {
		return proof, err
	}

	err = wire.ReadBinaryBytes(proof.Data(), data)

	return proof, err
}

func GetHeight() int {
	return viper.GetInt(heightFlag)
}

// GetProof performs the 'get' command directly from the proof.. (not from the CLI)
func GetProof(node client.Client, prover lite_client.Prover, key []byte, height int) (lite_client.Proof, error){
	proof, err := prover.Get(key, uint64(height))
	if err != nil {
		return nil, err
	}

	proofHeight := int64(proof.BlockHeight())
	// here is the certifier, root of all knowledge
	cert, err := commands.GetCertifier()

	// get and validate a signed header for this proof..
	client.WaitForHeight(node, proofHeight, nil)
	commit, err := node.Commit(&proofHeight)
	if err != nil {
		return nil,err
	}

	// commit include a majority of signatures from the last known validators set..
	// https://tendermint.readthedocs.io/projects/tools/en/master/specification/light-client-protocol.html?highlight=commit
	check := lite_client.Checkpoint{
		Header: commit.Header,
		Commit: commit.Commit,
	}

	// 검증...
	err = cert.Certify(check)
	if err != nil {
		return nil, err
	}

	// Validate the proof against the certified header to ensure data integrity
	err = proof.Validate(check)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%v %v", proof, err)	// to be removed
	return proof, err
}

type proof struct {
	Height 		uint64 			`json:"height"`
	Data		interface{}		`json:"data"`
}

func OutputProof( info interface{}, height uint64) error {

	wrap := proof{Height:height, Data:info}

	res, err := data.ToJSON(wrap)
	if err != nil {
		return err
	}
	fmt.Println(string(res))
	return nil

}