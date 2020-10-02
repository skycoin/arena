package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/skycoin/arena/pkg/adder"
)

func main() {
	a, b, err := parseArguments(os.Args[1:])
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
	adder.AddNumbers(a, b)
}

var notEnoughArgsErr = errors.New("Not enough arguments")

type parseError struct {
	value string
}

func (pe parseError) Error() string {
	return fmt.Sprintf("Invalid float value: %s", pe.value)
}

// parseArguments takes first two arguments and tries to parse them
// into floats
// return parsed values, or error:
// notEnoughArgsErr when there are less than two arguments
// parseFailedErr when one of the two
func parseArguments(args []string) (float64, float64, error) {
	var a, b float64
	if len(args) < 2 {
		return a, b, notEnoughArgsErr
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

func parseValue(val string) (float64, error) {
	panic("Not implemented")
}
