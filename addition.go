package main

import "fmt"

func main() {
	var firstNum, secondNum int

	fmt.Print("Enter first number - ")
	fmt.Scanln(&firstNum) //scan first number

	fmt.Print("Enter second number - ")
	fmt.Scanln(&secondNum) //scan second number

	total := firstNum + secondNum
	fmt.Print("The sum of given numbers is - ", total) // printing result as sum of given numbers
}
