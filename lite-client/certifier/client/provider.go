package client

import (
	"github.com/alexjipark/wabocoin/lite-client/certifier"
	rpcclient 	"github.com/tendermint/tendermint/rpc/client"
)

var _ certifier.Provider = &Provider{}

type Provider struct {
	node 		rpcclient.SignClient
	lastHeight 	int
}

func NewHTTP(remote string) *Provider {
	return &Provider{
		node: rpcclient.NewHTTP(remote, "/websocket"),
	}
}

func (p *Provider) StoreSeed (seed certifier.Seed) error {
	return nil
}

func (p *Provider) GetByHeight(h int) (certifier.Seed, error) {
	return certifier.Seed{}, nil
}

func (p *Provider) GetByHash(hash []byte) (certifier.Seed, error) {
	return certifier.Seed{}, nil
}
