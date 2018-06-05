package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tmlibs/cli"
	"github.com/pkg/errors"
)

type Runable func(cmd *cobra.Command, args []string) error


// Any commands that require and init'ed lite client directory
// should wrap RunE command with RequireInit to make sure that the client is initialized
//
// This cannot be called during PersistentRunE
// As they are called from the most specific command first, and root last..
// the root command sets up viper, which is needed to find the home dir..
func RequireInit(run Runable) Runable {
	return func(cmd *cobra.Command, args []string) error {
		// first check it was init'ed and if not, return an error
		root := viper.GetString(cli.HomeFlag)
		init, err := WasInited(root)
		if err != nil {
			return err
		}
		if !init {
			return errors.Errorf("you must run init first..")
		}

		return run(cmd, args)
	}
}


func WasInited(root string) (bool,error) {
	// root is directory..
	// file existence check..
	return true, nil
}
