package main

import (
	"arena/pkg"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	firstInt, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		err = handleErrors(err, 1)
		fmt.Println(err.Error())
	}

	secondInt, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		err = handleErrors(err, 2)
		fmt.Println(err.Error())
	}

	sum := pkg.Sum(firstInt, secondInt)
	fmt.Println(sum)
}

func handleErrors(err error, arg int) error {
	return fmt.Errorf("Error with argument %d, error: %w", arg, err)
}
