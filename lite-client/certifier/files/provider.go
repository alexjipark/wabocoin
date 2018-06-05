package files

import (
	"os"
	"path/filepath"
	"github.com/alexjipark/wabocoin/lite-client/certifier"
	"encoding/hex"
	"fmt"
)

const (
	Ext = ".tsd"
	ValDir = "validators"
	CheckDir = "checkpoints"
	dirPerm = os.FileMode(0755)
	filePerm = os.FileMode(0644)
)

// Provider stores all data in the file system..
// the Same validator hash may be reused by different headers/Checkpointers..
// validator hash - validator set..
// block height - Checkpoint with block height <= h
// Seed - store it..
type Provider struct {
	valDir 		string
	checkDir	string
}

// NewProvider creates the parent dir, sub dirs..
func NewProvider(dir string) Provider {
	valDir := filepath.Join(dir, ValDir)
	checkDir := filepath.Join(dir, CheckDir)

	for _,d := range []string{valDir, checkDir} {
		err := os.MkdirAll(d, dirPerm)
		if err != nil {
			panic(err)
		}
	}

	return Provider{valDir:valDir, checkDir:checkDir}
}

func (p Provider) encodeHash(hash []byte) string {
	return hex.EncodeToString(hash) + Ext
}

func (p Provider) encodeHeight(h int) string {
	return fmt.Sprintf("%012d%s", h, Ext)
}

func (p Provider) StoreSeed(seed certifier.Seed) error {
	err := seed.ValidateBasic(seed.Header.ChainID)
	if err != nil {
		return err
	}

	paths := [] string {
		filepath.Join(p.checkDir, p.encodeHeight(int(seed.Header.Height))),
		filepath.Join(p.valDir, p.encodeHash(seed.Header.ValidatorsHash)),
	}

	for _,path := range paths {
		err := seed.Write(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p Provider) GetByHeight(h int) (certifier.Seed, error) {

	path := filepath.Join(p.checkDir, p.encodeHeight(h))
	seed, err := certifier.LoadSeed(path)

	if certifier.IsSeedNotFoundError(err) {

	}

	return certifier.Seed{}, nil
}

func (m Provider) searchForHeight(h int) (string, error) {

}

func (p Provider) GetByHash(hash []byte) (certifier.Seed, error) {

	path := filepath.Join(p.valDir, p.encodeHash(hash))

	return certifier.LoadSeed(path)
}
