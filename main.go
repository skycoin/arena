package main

import (
	"fmt"
	"os"
	"strconv"
)


func handleInputError(err error)  {
	fmt.Printf("Error occurred reading input argument (int/float value expected): %s\n", err.Error())
}


func main()  {

	if len(os.Args) < 3 {
		fmt.Println("Not enough arguments: 2 int/float numbers required")
		return
	}

	a, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		handleInputError(err)
		return
	}

	b, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		handleInputError(err)
		return
	}

	fmt.Printf("%v\n", a + b)
}
