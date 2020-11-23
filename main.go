package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter first number: ")

	firstOperandInputBytes, _, readingErrorForFirstOperand := reader.ReadLine()
	if readingErrorForFirstOperand != nil {
		fmt.Fprintf(os.Stderr, "Error while parsing input! %v\n", readingErrorForFirstOperand)
		os.Exit(1)
	}
	firstOperandInputString := string(firstOperandInputBytes)

	fmt.Print("Enter second number: ")
	secondOperandInputBytes, _, readingErrorForSecondOperand := reader.ReadLine()

	if readingErrorForSecondOperand != nil {
		fmt.Fprintf(os.Stderr, "Error while parsing input! %v\n", readingErrorForSecondOperand)
		os.Exit(1)
	}
	secondOperandInputString := string(secondOperandInputBytes)

	firstOperand, parsingErrorForFirstOperand := strconv.ParseFloat(firstOperandInputString, 64)
	secondOperand, parsingErrorForSecondOperand := strconv.ParseFloat(secondOperandInputString, 64)

	if parsingErrorForFirstOperand != nil {
		fmt.Fprintf(os.Stderr, "Error while parsing number: %v\n", parsingErrorForFirstOperand)
		os.Exit(1)
	}
	if parsingErrorForSecondOperand != nil {
		fmt.Fprintf(os.Stderr, "Error while parsing number: %v\n", parsingErrorForSecondOperand)
		os.Exit(1)
	}
	fmt.Println("Sum of given numbers is:", Add(firstOperand, secondOperand))
}

func Add(firstOperand float64, secondOperand float64) float64 {
	return firstOperand + secondOperand
}
