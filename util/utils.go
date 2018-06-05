package util

import (
	"github.com/tendermint/tmlibs/common"
	"encoding/hex"
	"github.com/pkg/errors"
)

func ParseHexKey(key string) ([]byte, error) {
	r, err := hex.DecodeString(common.StripHex(key))
	return r, errors.WithStack(err)
}
