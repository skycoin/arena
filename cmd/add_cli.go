package main

import (
	"fmt"
	"os"
	"strings"
	"io"

	"github.com/mungujn/arena"
)

func main() {
	var stream io.Reader

	if parameters := os.Args; len(parameters) == 3 {
		stream = strings.NewReader(parameters[1] + " " + parameters[2])
	} else {
		fmt.Println("Input two numbers:")
		stream = os.Stdin
	}

	firstNumber, secondNumber, err := arena.ReadNumbers(stream)
	if err != nil {
		fmt.Println("Incorrect inputs")
		return
	}
	sum := arena.AddNumbers(firstNumber, secondNumber)
	fmt.Printf("Sum is: %v \n", sum)
}
