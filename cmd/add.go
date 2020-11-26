package cmd

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Writing to out through io.Writer to facilitate unit testing
var out io.Writer = os.Stdout

var errCannotConvert = fmt.Errorf("Cannot convert first argument to number")
var errTwoArguments = fmt.Errorf("Exactly two arguments required to add")

// AddCommand adds two numbers given as arguments
func AddCommand() {
	addCommand := flag.NewFlagSet("add", flag.ExitOnError)

	// Two arguments are required, exit if one of them is not provided
	if len(os.Args) < 2 {
		fmt.Fprintf(out, "Enter a valid command\n")
		addCommand.Usage()
		os.Exit(1)
	}

	if err := addCommand.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(out, "Cannot parse, Err: %v", err)
	}

	sum, err := add(addCommand.Args())
	if err != nil {
		fmt.Fprintf(out, "Cannot add, Err: %v", err)
	} else {
		fmt.Fprintln(out, sum)
	}
}

func add(args []string) (int, error) {
	if len(args) != 2 {
		return 0, errTwoArguments
	}

	firstNum, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, errCannotConvert
	}
	secondNum, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, errCannotConvert
	}

	return firstNum + secondNum, nil
}
