package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Please provide exactly 2 numbers")
		os.Exit(1)
	}

	add1, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Println("First argument is invalid, it is expected to be a number")
		os.Exit(1)
	}

	add2, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("Second argument is invalid, it is expected to be a number")
		os.Exit(1)
	}

	fmt.Println(add1 + add2)
}
