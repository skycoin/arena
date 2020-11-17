package adder

import (
	"errors"
	"fmt"
)

var (
	// ErrArgCount args count validation error
	ErrArgCount = errors.New("invalid args count")
)

// ArgError arg validation error
type ArgError struct {
	argIx int
	Cause error
}

// NewArgError constructor for ArgError
func NewArgError(argIx int, cause error) error {
	return &ArgError{argIx: argIx, Cause: cause}
}

func (a ArgError) Error() string {
	return fmt.Sprintf("invalid number%d cause: %s", a.argIx+1, a.Cause.Error())
}
