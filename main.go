package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const delimiter = " "

// SumStrings takes two strings num1 and num2
// It returns the sum of these two strings if they are both numeric.
// Otherwise it returns an error
func SumStrings(s1, s2 string) (float64, error) {
	var num1, num2 float64
	var err error

	if num1, err = strconv.ParseFloat(s1, 64); err != nil {
		return 0, errors.New("1st argument is not a number")
	}

	if num2, err = strconv.ParseFloat(s2, 64); err != nil {
		return 0, errors.New("2nd argument is not a number")
	}

	return num1 + num2, nil
}

func main() {
	args := os.Args[1:]
	var err error

	if len(args) != 2 {
		// read numbers from terminal
		fmt.Println("Arguments not set properly. Please input two numbers, separate by a space. For example: 1 2")

		reader := bufio.NewReader(os.Stdin)

		var bytes []byte

		if bytes, _, err = reader.ReadLine(); err != nil {
			fmt.Println("Failed reading input")
			return
		}

		args = strings.Split(string(bytes), delimiter)
		if len(args) != 2 {
			fmt.Println("Failed reading 2 numbers")
			return
		}
	}

	var result float64
	if result, err = SumStrings(args[0], args[1]); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
