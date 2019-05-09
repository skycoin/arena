package main

import (
	"fmt"
	"math/big"
	"os"

	"github.com/skycoin/arena/pkg/mathutil"
)

const (
	usage        = "Usage: add <num1> <num2>"
	bigFloatBase = 0
)

func main() {
	if len(os.Args) != 3 {
		_, _ = fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	arg1, _, err := new(big.Float).Parse(os.Args[1], bigFloatBase)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error parsing arg \"%s\": %v\n", os.Args[1], err)
		os.Exit(1)
	}

	arg2, _, err := new(big.Float).Parse(os.Args[2], bigFloatBase)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error parsing arg \"%s\": %v\n", os.Args[2], err)
		os.Exit(1)
	}

	sum := mathutil.Sum(arg1, arg2)

	fmt.Println(sum)
}
