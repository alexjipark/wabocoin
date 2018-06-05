package tx

import (
	"github.com/spf13/viper"
	"io"
	"os"
	"io/ioutil"
	"github.com/pkg/errors"
	"encoding/json"
	"github.com/tendermint/go-crypto"
	"github.com/alexjipark/wabocoin/cmd/wabocli/cmd/command"
	"github.com/alexjipark/wabocoin/lite-client/commands"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"

	)

func LoadJSON (template interface{}) (bool, error) {
	input := viper.GetString(InputFlag)
	if input == "" {
		return true, nil
	}

	raw, err := readInput(input)

	// parse the input
	err = json.Unmarshal(raw, template)
	if err != nil {
		return true, err
	}
	return true, nil


}


func readInput(file string) ([]byte, error) {
	var reader io.Reader
	if file == "-" {
		reader = os.Stdin
	} else {
		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		reader = f
	}

	data, err := ioutil.ReadAll(reader)
	return data, errors.WithStack(err)
}


// GetSigner will return the pub key that will sign the tx
// [AJ] How to sign with pub key??..
func GetSigner() crypto.PubKey {
	return nil
}

func Sign(tx interface{}) ([]byte, error) {

	return nil
}

// SignAdnPostTx does all work once we construct a proper struct
// it validates data , signs if needed,  transform into bytes
// and posts to the node..
func SignAndPostTx(tx *command.SendTx) (*ctypes.ResultBroadcastTxCommit, error) {

	err := tx.ValidateBasic()

	packet, err := Sign(tx)

	// post the bytes
	node := commands.GetNode()
	return node.BroadcastTxCommit(packet)

}








