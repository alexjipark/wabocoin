package certifier

import (
	rawerr "errors"
	"github.com/pkg/errors"
)

var (
	errSeedNotFound = rawerr.New("Seed Not Found by Provider")
	errValidatorChanged = rawerr.New("Validators differ from header and certifier")

)

func IsSeedNotFoundError(err error) bool {
	return err != nil && (errors.Cause(err) == errSeedNotFound)
}

func IsValidatorChangedErr (err error) bool {
	return err != nil && (errors.Cause(err) == errValidatorChanged)
}