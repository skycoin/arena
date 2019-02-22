package main

import (
	"fmt"
	"strconv"
)

func main() {
	var result float64
	for i := 0; i < 2; {
		fmt.Print("Enter number: ")
		var text string
		_, err := fmt.Scanln(&text)
		if err != nil {
			continue
		}

		number, err := strconv.ParseFloat(text, 64)
		if err != nil {
			fmt.Printf("%v - is not a number\n", text)
			continue
		}

		result += number
		i++
	}

	fmt.Printf("Finish sum: %v\n", result)
}
