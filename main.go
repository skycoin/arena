package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sum float64

	var input string

	i := 2
	for i > 0 {
		fmt.Println("Enter number:")
		fmt.Scanln(&input)
		tmp, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Printf("Wrong input! {%s}\n", err.Error())
		} else {
			sum += tmp
			i--
		}
	}
	fmt.Printf("Result: %v\n", sum)
}
