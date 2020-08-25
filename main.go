package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Please provide 2 numbers for addition")
		os.Exit(1)
	}

	n, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Println("First argument is not a number, it is expected to be a number")
		os.Exit(1)
	}

	m, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("Second argument is not a number, it is expected to be a number")
		os.Exit(1)
	}

	fmt.Println("Answer is: ", n+m)
}
