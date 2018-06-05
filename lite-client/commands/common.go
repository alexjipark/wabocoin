package commands

import (
	rpcClient "github.com/tendermint/tendermint/rpc/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/alexjipark/wabocoin/lite-client/certifier"
	"errors"
	"github.com/tendermint/tmlibs/cli"
	"github.com/alexjipark/wabocoin/lite-client/certifier/client"
	"github.com/alexjipark/wabocoin/lite-client/certifier/files"
)

var (
	trustedProv		certifier.Provider
	sourceProv 		certifier.Provider
)

const (
	ChainFlag = "chain-id"
	NodeFlag = "node"
)

func AddBasicFlags (cmd *cobra.Command) {
	cmd.PersistentFlags().String(ChainFlag, "", "Chain Id of Tendermint Node")
	cmd.PersistentFlags().String(NodeFlag, "", "<host>:<port> to tendermint rpc interface for this chain..")
}

func GetChainID() string {
	return viper.GetString(ChainFlag)
}

func GetNode() rpcClient.Client {
	return rpcClient.NewHTTP(viper.GetString(NodeFlag), "/websocket")	//[AJ?] websocket means??
}

func GetProviders() (trusted certifier.Provider, source certifier.Provider) {
	if trustedProv == nil || sourceProv == nil {
		rootDir := viper.GetString(cli.HomeFlag)
		trustedProv = certifier.NewCacheProvider(
			certifier.NewMemStoreProvider(),
			files.NewProvider(rootDir),
		)
		node := viper.GetString(NodeFlag)
		sourceProv = client.NewHTTP(node)
	}
	return trustedProv, sourceProv
}

func GetCertifier() (*certifier.InquiringCertifier, error) {
	// load up the latest store..
	trusted, source := GetProviders()

	// gets the most recent verified one..
	seed, err := certifier.LatestSeed(trusted)

	if certifier.IsSeedNotFoundError(err) {
		return nil, errors.New("Please run init first to establish a root of trust")
	}
	if err != nil {
		return nil, err
	}
	cert := certifier.NewInquiring(GetChainID(), seed.Validators, trusted, source)
	return cert, nil
}

