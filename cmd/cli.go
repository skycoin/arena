package cmd

import (
	"bufio"
	"fmt"
	"math/big"
	"os"

	"github.com/arena/src/add"
	"github.com/arena/src/utils"
)

const maxAttempts = 3

// Cmd is to run command line program
type Cmd struct{}

// Add will run the addition program
func (c *Cmd) Add() {
	// Initialize data object
	data := add.Data{}

	//Enter first number
	fmt.Print("Enter first number: ")
	data.Numbers = append(data.Numbers, fetchNumber(maxAttempts))

	// Enter second number
	fmt.Print("Enter second number: ")
	data.Numbers = append(data.Numbers, fetchNumber(maxAttempts))
	// We can extend this program to add more number if required.

	fmt.Println("The final result is: ", data.Add())
}

// Fetch number and validate
func fetchNumber(try int) *big.Float {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	number, err := utils.ParseBigFloat(scanner.Text())
	if err != nil {
		if try == 1 {
			fmt.Println("Max attempts reached, Closing program !!!")
			os.Exit(1)
		}
		fmt.Printf("Invalid format, Enter a number again: ")
		return fetchNumber(try - 1)
	}
	return number
}
