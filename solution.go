package main

import (
	"fmt"
)

func main() {
	fmt.Println("This program find sum and print her in standart output")

	first := scanNumber("Enter first number")
	second := scanNumber("Enter second number")

	fmt.Println("sum: ", first+second)
}

func scanNumber(text string) int {
	var number int

	for {
		fmt.Println(text)
		_, err := fmt.Scanf("%d\n", &number)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		break
	}
	return number
}
