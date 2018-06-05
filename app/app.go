package app

import (
	dbm "github.com/tendermint/tmlibs/db"
	"encoding/json"
	"github.com/tendermint/abci/types"

	"github.com/alexjipark/wabocoin/code"
	wt "github.com/alexjipark/wabocoin/types"
	"github.com/tendermint/go-wire"
)

type WaboState struct {
	// Database
	DB 			dbm.DB
	// State - Size, Block Height, AppHash
	Size		uint64
	Height 		uint64
	AppHash		[]byte

}

var (
	stateKey = []byte("stateKey")
)

const (
	maxTxSize 	= 10240
)

// Load State
func loadState (db dbm.DB) WaboState {
	stateBytes := db.Get(stateKey)
	var state WaboState

	if len(stateBytes) != 0 {
		err := json.Unmarshal( stateBytes, &state)
		if err != nil {
			panic(err)
		}
	}
	state.DB = db
	return state
}

// Save State
func saveState (state WaboState) {
	stateBytes, err := json.Marshal(state)
	if err != nil {
		panic(err)
	}

	state.DB.Set( stateKey, stateBytes)
}

var _ types.Application = (*WabocoinApp)(nil)

type WabocoinApp struct {
	types.BaseApplication

	state WaboState
}

func NewWabocoinApp() *WabocoinApp {
	state := loadState( dbm.NewMemDB() )
	return &WabocoinApp{ state:state}
}

func (app *WabocoinApp) Info (req types.RequestInfo) (resInfo types.ResponseInfo) {
	return
}

func (app *WabocoinApp) SetOption (req types.RequestSetOption) (resSO types.ResponseSetOption) {
	return
}

func (app *WabocoinApp) DeliverTx (txBytes []byte) (res types.ResponseDeliverTx) {
	if len(txBytes) > maxTxSize {
		return types.ResponseDeliverTx{Code:code.CodeErrBaseEncodingError}
	}

	// Decode Tx..
	var tx wt.Tx
	err := wire.ReadBinaryBytes(txBytes, &tx)
	if err != nil {
		return types.ResponseDeliverTx{Code:code.CodeErrBaseEncodingError}
	}

	return
}

func (app *WabocoinApp) CheckTx (tx []byte) types.ResponseCheckTx {

	return types.ResponseCheckTx{ Code: code.CodeTypeOK }
}

func (app *WabocoinApp) Commit() (res types.ResponseCommit) {
	return
}

func (app *WabocoinApp) Query (req types.RequestQuery)  (res types.ResponseQuery) {
	return
}

func (app *WabocoinApp) InitChain (req types.RequestInitChain) (res types.ResponseInitChain) {
	return
}

func (app *WabocoinApp) BeginBlock (req types.RequestBeginBlock) (res types.ResponseBeginBlock) {
	return
}

func (app *WabocoinApp) EndBlock (req types.RequestEndBlock) (res types.ResponseEndBlock) {
	return
}














