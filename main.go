package main

import "fmt"

func main() {
	prnt := fmt.Printf
	var firstNmb float64
	var secondNmb float64

	prnt("Adding two numbers\n")
	prnt("enter first number: ")
	fmt.Scanf("%f", &firstNmb)
	prnt("enter second number:")
	fmt.Scanf("%f", &secondNmb)
	prnt("%.2f + %.2f = %.2f\n", firstNmb, secondNmb, firstNmb+secondNmb)
}
