package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/skycoin/Addition/pkg/add"
)

const maxAttempts = 3

func main() {

	data := add.Data{}
	fmt.Print("Enter first number: ")
	data.Numbers = append(data.Numbers, takeNumber(maxAttempts))
	fmt.Print("Enter second number: ")
	data.Numbers = append(data.Numbers, takeNumber(maxAttempts))

	fmt.Println("The final result is: ", data.Add())
}

func takeNumber(try int) float64 {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	number, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		if try == 0 {
			fmt.Println("Max attempts reached, Closing program !!!")
			os.Exit(1)
		}
		fmt.Printf("Invalid format, Enter a number again: ")
		return takeNumber(try - 1)
	}
	return number
}
