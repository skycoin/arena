package main

import (
	"fmt"
)

func main() {

	var numOne, numTwo float32

	fmt.Println("Welcome. Please enter first number to sum: ")
	fmt.Scanln(&numOne)

	fmt.Println("Please enter second number to sum: ")
	fmt.Scanln(&numTwo)

	sum := numTwo + numOne
	fmt.Println("The sum of the numbers is ", sum)
	fmt.Println("Goodbye!")

}
