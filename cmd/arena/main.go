package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			str := r.(string)
			fmt.Println(str)
		}
	}()

	args := os.Args[1:]
	if len(args) != 2 {
		panic("please provide exactly 2 parameters")
	}

	num1, err := strconv.ParseFloat(args[0], 4)
	if err != nil {
		panic("Can't convert 1st argument to float. Error: " + err.Error())
	}

	num2, err := strconv.ParseFloat(args[1], 4)
	if err != nil {
		panic("Can't convert 2nd argument to float. Error: " + err.Error())
	}

	fmt.Println("Sum: ", num1+num2)
}
