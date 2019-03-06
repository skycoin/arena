package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readNumber(message string) float64 {
	fmt.Printf(message)

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		if scanner.Text() == "exit" {
			os.Exit(0)
		}

		result, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			println("Enter correct number:")
		} else {
			return result
		}
	}
}

func main() {

	firstNumber := readNumber("Enter first number or type 'exit':")
	secondNumber := readNumber("Enter second number or type 'exit':")
	fmt.Printf("Result is %f", firstNumber+secondNumber)

}

