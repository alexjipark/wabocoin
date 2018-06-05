package lite_client

import (
	"github.com/pkg/errors"
	"fmt"
)

type errHeightMismatch struct {
	h1, h2 int
}

func (e errHeightMismatch) Error() string {
	return fmt.Sprintf("Block don't match - %d vs %d", e.h1, e.h2)
}

func IsHeightMismatchError(err error) bool {
	if err == nil {
		return false
	}
	_, ok := errors.Cause(err).(errHeightMismatch)
	return ok
}

func ErrHeightMismatch(h1, h2 int) error {
	err := errHeightMismatch{h1,h2}
	return errors.WithStack(err)
}

//============//

type  errNoData struct{}

func (e errNoData) Error() string {
	return fmt.Sprintf("No Data returned for a query..")
}

// IsNoDataErr checks whether an error is due to a query returning empty data
func IsNoDataErr(err error) bool {
	if err == nil {
		return false
	}

	_, ok := errors.Cause(err).(errNoData)	// [AJ?] coding rule checks!
	return ok
}


func ErrNoData() error {
	return errors.WithStack(errNoData{})
}