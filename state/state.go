package state

import "github.com/cosmos/cosmos-sdk/types"

type State struct {
	chainID		string
	store 		types.KVStore
}

func (s *State) GetChainID() string {
	if s.chainID != "" {
		return s.chainID
	}
	s.chainID = string(s.store.Get([]byte("base/chain_id")))
	return s.chainID
}