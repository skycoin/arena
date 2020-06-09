package main

import "fmt"

func main() {
	fmt.Print("Enter First number ")
	var num1 int
	fmt.Scanln(&num1)
	fmt.Print("Enter Second number ")
	var num2 int
	fmt.Scanln(&num2)
	var add = num1 + num2
	fmt.Print(add)
}
