package main

/*
# Task

Write a Golang program
- that takes two numbers from the command line,
- adds them up
- and prints the result.
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
		result += value
		if err != nil {
			return fmt.Sprintf("%v", err)
		}
	}
	return fmt.Sprintf("%v", result)
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	fmt.Println(sumargs(os.Args[1], os.Args[2]))
}
