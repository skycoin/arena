package main

import (
	"fmt"
)

func main() {
	var x1, x2 int
	fmt.Printf("Please, enter two numbers.\ninput x1: ")
	fmt.Scan(&x1)

	fmt.Print("input x2: ")
	fmt.Scan(&x2)

	sum := x1 + x2

	fmt.Println("result:", sum)
}
