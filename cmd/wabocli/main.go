package main

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"

	//"github.com/tendermint/light-client/commands"
	wcmd "github.com/alexjipark/wabocoin/cmd/wabocli/cmd"

)

var RootCmd = &cobra.Command{
	Use: "wabocli",
	Short: "Light Client for Wabocoin",
	Run: wabo_cli,
}

func init() {
	RootCmd.PersistentFlags().String("lang", "en", "language to Use")
	RootCmd.AddCommand(wcmd.GreetPlanetGmd)
}

func main() {

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
/*
	commands.AddBasicFlags(RootCmd)	// chain-id (string) , node (rpc interface - <host>:<port>)

	// Prepare Queries...
	pr := proofs.RootCmd			// 'query'
	pr.AddCommand(proofs.TxCmd)		// 'tx [txhash]'
	pr.AddCommand(proofs.KeyCmd)	// 'key [key]'
*/


}

func wabo_cli (cmd *cobra.Command, args []string) {

	lang := cmd.Flag("lang").Value.String()
	greeting := wcmd.Greeting(lang)

	if len(args) == 0 {
		fmt.Printf("%s travellers!!\n", greeting)
		os.Exit(1)
	}

	fmt.Printf("wabo_cli in %s\n", greeting )


}
