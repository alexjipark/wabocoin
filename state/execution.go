package state

import (
	"github.com/alexjipark/wabocoin/types"
	abci "github.com/tendermint/abci/types"
	"github.com/alexjipark/wabocoin/code"
	"github.com/tendermint/tmlibs/common"
)

func ExecTx (state *State, tx types.Tx) (res abci.ResponseDeliverTx) {

	chainID := state.GetChainID()
	switch tx := tx.(type) {

	case *types.SendTx:
		common.Fmt("info : %v %v", chainID, tx.Gas)
		return

	case *types.AppTx:
		common.Fmt("info : %v %v", chainID, tx.Name)
		return
	}



	return abci.ResponseDeliverTx{Code:code.CodeTypeOK}
}

