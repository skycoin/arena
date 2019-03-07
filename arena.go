package main

/*
# Task

Write a Golang program:
 - that takes two numbers from the command line,
 - adds them up
 - and prints the result

# Analysis

- "two numbers"
- exact number type was not defined
- it's not defined what to do with parsing errors and overflows, Infinities and NaN
- format of print not specified

# Plan
 - Create `dep` project
 - Write a program that sums exactly *two* numbers
 - Use float64 for parsing and summing values
 - Make `main` as simple as possible: check number of `os.Args` and `fmt.Println(sumargs ...)`
 - `sumargs` - function of two `string` args, returns `string` always.
 	In case of ParseFloat errors returns error message
 - Create table-driven tests for `sumargs`. Include test cases for ParseFloat errors
 - Create Makefile as described in Idiomatic Go (in Skycoin)
 - Make comments where appropriate in source files
 - Check guidelines for compliance
 - Commit
 - PR
*/

import (
	"fmt"
	"os"
	"strconv"
)

func sumargs(argx, argy string) string {
	result := 0.0
	for _, arg := range []string{argx, argy} {
		value, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			return fmt.Sprintf("%v", err)
		}
		result += value
	}
	return fmt.Sprintf("%v", result)
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	fmt.Println(sumargs(os.Args[1], os.Args[2]))
}
