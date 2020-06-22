// Program takes two numbers from the command line, adds them up and prints the result
package main

import (
	"fmt"
	"os"
)

func main() {
	num1, num2, res, err := MakeOperation(os.Stdin)
	if err != nil {
		fmt.Printf("\n An error occurred: %v", err)
		return
	}
	fmt.Printf("\nResult of additing two numbers %f and %f are: %f", num1, num2, res)
}
