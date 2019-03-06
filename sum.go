package main

import (
	"fmt"
	"os"
	"strconv"
)

func sumDigits(dig1 int, dig2 int) int {
	return dig1 + dig2
}

func main() {
	// check whether the count and validaty of arguments
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <digit1> <digit2>\n", os.Args[0])
		os.Exit(1)
	}

	dig1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Invalid integer value '%s'\n", os.Args[1])
		os.Exit(1)
	}

	dig2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid integer value '%s'\n", os.Args[2])
		os.Exit(1)
	}

	fmt.Printf("The sum is %d\n", sumDigits(dig1, dig2))
}
