package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	argsWithoutProg := os.Args[1:]
	first, err := strconv.ParseFloat(argsWithoutProg[0], 64)
	if err != nil {
		fmt.Println("Please, input a number, your first input incorrect:", err)
		return
	}
	second, err := strconv.ParseFloat(argsWithoutProg[1], 64)
	if err != nil {
		log.Println("Please, input a number, your second input incorrect:", err)
		return
	}
	sum := Sum(first, second)
	fmt.Println(sum)
}
