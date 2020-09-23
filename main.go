package main

import (
	"flag"
	"fmt"
)

func main() {

	n1 := flag.Int("num1", 0, "must be an int")
	n2 := flag.Int("num2", 0, "must be an int")

	flag.Parse()

	result := *n1 + *n2

	fmt.Println("Sum of number is: ", result)
}
