package command

import (
	"github.com/spf13/cobra"
	"github.com/alexjipark/wabocoin/lite-client/commands"
	"github.com/alexjipark/wabocoin/types"
	tx2 "github.com/alexjipark/wabocoin/lite-client/commands/tx"
	"github.com/spf13/viper"
)

const (
	FlagTo = "to"
	FlagAmount = "amount"
	FlagFee = "fee"
	FlagGas = "gas"
	FlagSequence = "sequence"
)


var SendTxCmd = &cobra.Command{
	Use: "send",
	Short: "send token from one account to another",
	RunE: commands.RequireInit(doSendTx),
}

func init() {
	flags := SendTxCmd.Flags()
	flags.String(FlagTo, "","Destination for the bits")
	flags.String(FlagAmount, "", "Coins to send in the format <amount><coin>,<amount><coin>...")
	flags.String(FlagFee, "0mycoin", "")
	flags.Int64(FlagGas, 0, "")
	flags.Int(FlagSequence, -1, "sequence number for the transaction..")

}

func doSendTx (cmd *cobra.Command, args []string) error {

	tx := new(types.SendTx)
	found, err := tx2.LoadJSON(tx)
	if err != nil {
		return err
	}
	if !found {

	}

	err = readSendTxFlag(tx)
	if err != nil {
		return err
	}

	//====
	send := &SendTx{
		chainID: commands.GetChainID(),
		Tx: tx,
	}

	send.AddSigner(tx2.GetSigner())

	




	return nil
}

func readSendTxFlag(tx *types.SendTx) error {

	var to []byte
	viper.GetString(FlagTo)

	// Set Gas, Fee, Inputs, Outputs in SendTx.. // Parsing!!!
	fees, err := tx2.ParseCoin(viper.GetString(FlagFee))

	amountCoins, err := tx2.ParseCoins(viper.GetString(FlagAmount))
	if err != nil {
		return err
	}

	tx.Gas = viper.GetInt64(FlagGas)

	tx.Inputs = []types.TxInput{{
		Coins: amountCoins,
		Sequence: viper.GetInt(FlagSequence),
	}}

	tx.Outputs = []types.TxOutput {{
		Address: to,
		Coins: amountCoins,
	}}



	viper.GetString(FlagSequence)
	return nil
}