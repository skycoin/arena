package adder

import (
	"errors"
	"fmt"
	"math/big"
)

const precision = 100

var NotEnoughArgsErr = errors.New("Not enough arguments")

type ParseError struct {
	value string
}

func (pe ParseError) Error() string {
	return fmt.Sprintf("Invalid float value: %s", pe.value)
}

// ParseArguments takes first two arguments and tries to parse them
// into floats
// return parsed values, or error:
// notEnoughArgsErr when there are less than two arguments
// parseFailedErr when one of the two
func ParseArguments(args []string) (*big.Float, *big.Float, error) {
	var a, b *big.Float
	if len(args) < 2 {
		return a, b, NotEnoughArgsErr
	}
	// use loop and return a slice if the number of arguments to sum becomes higher
	a, err := parseValue(args[0])
	if err != nil {
		return a, b, err
	}
	b, err = parseValue(args[1])
	if err != nil {
		return a, b, err
	}
	return a, b, nil
}

func parseValue(s string) (*big.Float, error) {
	res, _, err := big.ParseFloat(s, 0, precision, big.ToNearestEven)
	return res, err
}

func AddNumbers(a, b *big.Float) *big.Float {
	panic("Not implemented")
}
