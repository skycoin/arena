package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter First number: ")
	var first float64
	fmt.Scanln(&first)
	fmt.Print("Enter Second number: ")
	var second float64
	fmt.Scanln(&second)
	fmt.Println(first + second)
	fmt.Print("to 2 decimal places: ")
	fmt.Printf("%.2f", first+second)

}
