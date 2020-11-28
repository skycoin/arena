package main

import (
	"fmt"
	"math/big"
	"os"

	"github.com/Mikelle/arena/pkg/service"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %v <num1> <num2>\n", os.Args[0])
		os.Exit(0)
	}

	num1, _, err := new(big.Float).Parse(os.Args[1], 0)
	if err != nil {
		fmt.Printf("can't parse %v: %v\n", os.Args[1], err)
		os.Exit(0)
	}

	num2, _, err := new(big.Float).Parse(os.Args[2], 0)
	if err != nil {
		fmt.Printf("can't parse %v: %v\n", os.Args[1], err)
		os.Exit(0)
	}

	result := service.SumTwoNumbers(num1, num2)

	fmt.Printf("%v\n", result)
}
