package main

import (
	"fmt"
	"os"
	"strconv"
)

func printUsage() {
	fmt.Printf("Usage: ./add num1 num2\n")
}

func parseNums() (float64, float64, error) {
	first, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		return 0, 0, err
	}

	second, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		return 0, 0, err
	}

	return first, second, nil
}

func add(num1, num2 float64) float64 {
	return num1 + num2
}

func main() {
	if len(os.Args) != 3 {
		printUsage()
		os.Exit(1)
	}

	num1, num2, err := parseNums()
	if err != nil {
		fmt.Println("Error parsing numbers:", err)
		os.Exit(1)
	}

	fmt.Println(add(num1, num2))
}
