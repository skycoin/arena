package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func add(args []string) (float64, error) {
	if len(args) < 2 {
		return 1, errors.New("Please enter two numbers separated by space")
	}

	num1, err1 := strconv.ParseFloat(args[0], 64)
	num2, err2 := strconv.ParseFloat(args[1], 64)

	if err1 != nil {
		return 2, errors.New("First argument must be a number")
	}

	if err2 != nil {
		return 2, errors.New("Second argument must be a number")
	}

	return num1 + num2, nil
}

func main() {
	result, err := add(os.Args[1:])

	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(int(result))
	}
	if result == float64(int64(result)) {
		fmt.Printf("Sum of %s and %s is : %.0f \n", os.Args[1], os.Args[2], result)
	} else {
		fmt.Printf("Sum of %s and %s is : %.2f \n", os.Args[1], os.Args[2], result)
	}
}
