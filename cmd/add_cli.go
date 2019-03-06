package main

import (
	"fmt"
	"os"

	"github.com/mungujn/arena"
)

func main() {
	fmt.Println("Input two numbers:")
	firstNumber, secondNumber, err := arena.ReadNumbers(os.Stdin)
	if err != nil {
		fmt.Println("Incorrect inputs")
		return
	}
	sum := arena.AddNumbers(firstNumber, secondNumber)
	fmt.Printf("Sum is: %v \n", sum)
}
