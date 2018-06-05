package command

import (
	"github.com/spf13/cobra"
	"github.com/alexjipark/wabocoin/util"
	"github.com/alexjipark/wabocoin/types"
	proof2 "github.com/alexjipark/wabocoin/lite-client/commands/proof"
	"github.com/alexjipark/wabocoin/lite-client"
	"github.com/pkg/errors"
)

var AccountQueryCmd = &cobra.Command{
	Use: "account [address]",
	Short: "Get details of an account, with Proof!!",
	RunE: doAccountQuery,	// RunE : returns Error..

}

func doAccountQuery (cmd *cobra.Command, args []string) error {

	addr, err := util.ParseHexKey(args[0])
	if err != nil {
		return err
	}
	key := types.AccountKey(addr)	// wabo/a/{addr}

	acc := new(types.Account)

	proof, err := proof2.GetAndParseAppProof(key, &acc)

	if lite_client.IsNoDataErr(err) {
		return errors.Errorf("Account bytes are empty for address %X", addr)
	} else if err != nil {
		return err
	}

	return proof2.OutputProof(acc, proof.BlockHeight())
}
