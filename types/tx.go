package types

import (
	"github.com/tendermint/go-wire/data"
	"github.com/tendermint/tmlibs/common"
	"encoding/json"
	"github.com/tendermint/go-crypto"
)

type Tx interface {
	AssertIsTx()
	SignBytes(chainID string) []byte
}

// Types of Tx implementations
const (
	// Account Transactions
	TxTypeSend = byte(0x01)
	TxTypeApp  = byte(0x02)
	TxNameSend = "send"
	TxNameApp  = "app"
)

var txMapper data.Mapper		// alexjipark ??!!

func (_ *SendTx) AssertIsTx() {}
func (_ *AppTx)  AssertIsTx() {}

// Register both private key types with go-wire/data (and thus go-wire)
func init() {
	txMapper = data.NewMapper(TxS{}).
		RegisterImplementation(&SendTx{}, TxNameApp, TxTypeApp).
			RegisterImplementation(&AppTx{}, TxNameSend, TxTypeSend)
}

//TxS add json serialization to Tx
type TxS struct {
	Tx `json:"unwrap"`
}

func (p TxS) MarshalJson() ([]byte, error) {
	return txMapper.ToJSON(p.Tx)
}

func (p *TxS) UnmarshalJson(data []byte) (err error) {
	parsed, err := txMapper.FromJSON(data)
	if err == nil {
		p.Tx = parsed.(Tx)		// ???? alexjipark
	}
	return
}

//-------------------- SendTx ------------------//
// : Send coins to an address
type SendTx struct {
	Gas		int64		`json:"gas"`	// Gas
	// Input.. Output..
	Fee 	Coin 		`json:"fee"`	// Fee
	Inputs 	[]TxInput	`json:"input"`	// Input
	Outputs []TxOutput  `json:"output"`
}

func (tx *SendTx) SignBytes (chainID string) []byte {
	return nil
}

func (tx *SendTx) String() string {
	return common.Fmt("SendTx{%v}", tx.Gas)
}


//-------------------- AppTx -------------------//
// : Send a msg to a contract that runs in the vm..
type AppTx struct {
	Gas 	int64 				`json:"gas"`
	Name	string 				`json:"type"`
	Data	json.RawMessage		`json:"data"`

	//...
	Input 	TxInput				`json::"input"`

}

func (tx *AppTx) SignBytes (chainID string) []byte {
	return nil
}

func (tx *AppTx) String() string {
	return common.Fmt("AppTx {%v %v %v}", tx.Gas, tx.Name, tx.Data)
}

func (tx *AppTx) SetSignature(sig crypto.Signature) bool {
	tx.Input.Signature = sig
	return true
}


//-----------

type TxInput struct {
	Address   data.Bytes       `json:"address"`   // Hash of the PubKey
	Coins     Coins            `json:"coins"`     //
	Sequence  int              `json:"sequence"`  // Must be 1 greater than the last committed TxInput
	Signature crypto.Signature `json:"signature"` // Depends on the PubKey type and the whole Tx
	PubKey    crypto.PubKey    `json:"pub_key"`   // Is present iff Sequence == 0
}

//------------
type TxOutput struct {
	Address 	data.Bytes		`json:"address"`	// Hash of the public key
	Coins 		Coins			`json:"coins"`

}









