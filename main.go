package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum, err := AddNumbers()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sum of two numbers is: ", sum)
	}
}

func AddNumbers() (float64, error) {
	args := os.Args
	if len(args) < 3 {
		return 0, errors.New("provide two numbers")
	}
	if len(args) > 3 {
		return 0, errors.New("provide only two numbers")
	}
	num1, err := strconv.ParseFloat(args[1], 64)
	num2, err := strconv.ParseFloat(args[2], 64)

	if err != nil {
		return 0, errors.New("arguments should be numbers not string")
	}

	return num1 + num2, nil

}
