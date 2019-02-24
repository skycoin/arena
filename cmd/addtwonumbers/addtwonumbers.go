package main

import (
	"fmt"
	"math/big"
	"os"

	"github.com/skycoin/arena/pkg/addtwonumbers"
)

const (
	// Let math/big detect base.
	base = 0
)

func main() {
	if len(os.Args) != 3 {
		_, _ = fmt.Fprintf(os.Stderr, "Usage: %v <num1> <num2>\n", os.Args[0])
		os.Exit(1)
	}

	num1, _, err := new(big.Float).Parse(os.Args[1], base)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Could not parse %v: %v\n", os.Args[1], err)
		os.Exit(1)
	}

	num2, _, err := new(big.Float).Parse(os.Args[2], base)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Could not parse %v: %v\n", os.Args[2], err)
		os.Exit(1)
	}

	result := addtwonumbers.AddTwoNumbers(num1, num2)

	fmt.Printf("%v\n", result)
}
