package main

import (
	"fmt"
)

func main() {

	a := 0
	b := a
	fmt.Println("Addition of numbers")
	fmt.Print("Enter first number:")
	fmt.Scanln(&a)
	fmt.Print("Enter second number:")
	fmt.Scanln(&b)

	fmt.Printf("Addition is :%v", a+b)

}
