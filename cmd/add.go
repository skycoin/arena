package cmd

import (
	"flag"
	"fmt"
	"os"
)

// AddCommand adds two numbers given as arguments
func AddCommand() {
	addCommand := flag.NewFlagSet("add", flag.ExitOnError)

	var firstNum, secondNum int

	// Inbuilt flag package does not support positional arguments, so arguments are named
	addCommand.IntVar(&firstNum, "first", 0, "First Number to add. (Required)")
	addCommand.IntVar(&secondNum, "second", 0, "Second Number to add. (Required)")

	// Two arguments are required, exit if one of them is not provided
	if len(os.Args) < 2 {
		fmt.Printf("Enter required arguments\n")
		addCommand.Usage()
		os.Exit(1)
	}

	if err := addCommand.Parse(os.Args[2:]); err != nil {
		fmt.Printf("Cannot parse, %v", err)
	}

	fmt.Printf("%v + %v = %v\n", firstNum, secondNum, firstNum+secondNum)
}
