package certifier

import (
	"github.com/alexjipark/wabocoin/lite-client"
	"github.com/tendermint/tendermint/types"
	"math"
	"os"
	"github.com/tendermint/go-wire"
	"github.com/pkg/errors"
)

var FutureHeight = (math.MaxInt32 - 5)	// [AJ?]

// Seed is a checkpoint and the actual validator set
// the base info you need to update to a given point, assuming knowledge of previous validator set
type Seed struct {

	lite_client.Checkpoint				`json:"checkpoint"`
	Validators 	*types.ValidatorSet		`json:"validator_set"`

}

type Seeds []Seed

// Provider is used to get more validators by other means [AJ?]
type Provider interface {
	StoreSeed(seed Seed) error
	GetByHeight(h int)	(Seed, error)		// Returning closest seed at with height <= h
	GetByHash(hash []byte) (Seed, error)	// Returning a seed exactly matching this validator hash
}

func LatestSeed (p Provider) (Seed, error) {
	return p.GetByHeight(FutureHeight)
}
/*
	var f *os.File
	f, err = os.Create(path)
	if err == nil {
		var n int
		wire.WriteBinary(s, f, &n, &err)
		f.Close()
	}
	// we don't write, but this is not an error
	if os.IsExist(err) {
		return nil
	}
	return errors.WithStack(err)
 */
func (s Seed) Write (path string) (err error){
	var f *os.File
	f, err = os.Create(path)
	if err == nil {
		var n int
		wire.WriteBinary(s, f, &n, &err)
		f.Close()
	}
	if os.IsExist(err) {
		return nil
	}
	return errors.WithStack(err)
}

func LoadSeed(path string) (Seed, error) {
	return Seed{},nil
}


// CacheProvider allows you to place one or more caches in front of a source Provider
// It run through them in order until a match is found.
// So you can keep a local cache and check with the network if no data is there..
type CacheProvider struct {
	Providers []Provider
}

// [AJ?] the meaning of ...??
func NewCacheProvider(providers ...Provider) CacheProvider {
	return CacheProvider{		// [AJ?] Why not with pointer??
		Providers: providers,
	}
}
// [AJ?] "p *CacheProvider" doesn't work..
func (c CacheProvider) StoreSeed(seed Seed) error {

	for _,p := range c.Providers {
		err := p.StoreSeed(seed)
		if err != nil {
			break;
		}
	}

	return nil
}
func (c CacheProvider) GetByHeight(h int) (s Seed, err error) {

	for _,p := range c.Providers {
		ts, err := p.GetByHeight(h)
		if err == nil {
			if ts.Header.Height > s.Header.Height {
				s = ts
			}
			if ts.Header.Height == int64(h) {
				break
			}
		}
	}
	return s, err
}

func (c CacheProvider) GetByHash(hash []byte) (s Seed, err error) {

	for _,p := range c.Providers {
		s, err := p.GetByHash(hash)
		if err == nil {
			break
		}
	}

	return Seed{}, nil
}
