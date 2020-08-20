package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("you must pass two arguments")
		os.Exit(1)
	}

	a, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Println("can't first argument convert to float")
		os.Exit(1)
	}

	b, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("can't second argument convert to float")
		os.Exit(1)
	}

	fmt.Println(a + b)
}
