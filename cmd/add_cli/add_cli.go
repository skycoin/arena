package main

import (
	"fmt"
	"os"

	"github.com/atang152/arena/pkg/services"
)

func main() {

	const numToRead = 2

	reader := os.Stdin
	fmt.Printf("Please input %v numbers:\n", numToRead)
	numbers, err := services.ReadNumbers(reader, numToRead)

	if err != nil {
		fmt.Println("Incorrect inputs", err)
		return
	}

	sum := services.AddNumbers(numbers)
	fmt.Printf("The sum is %v \n", sum)

}
