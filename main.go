package main

import (
	"fmt"
	"os"
	"strconv"

	"tsocial.buri/pkg/errors"
)

func main() {
	numbers := os.Args
	var errInvalidArgument = errors.New(1, "arguments invalid")
	if len(numbers) != 3 {
		fmt.Printf("%q", errInvalidArgument)
		os.Exit(1)
	}
	leftnum, errLeft := strconv.Atoi(numbers[1])
	rightnum, errRight := strconv.Atoi(numbers[2])
	if errLeft != nil || errRight != nil {
		fmt.Printf("%q", errInvalidArgument)
		os.Exit(1)
	}
	fmt.Println(leftnum + rightnum)
}

// Contacts:
// olegmusinem@gmail.com
// @musinit
