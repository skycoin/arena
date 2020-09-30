package main

import (
	"arena/adder"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Not enough arguments.")
	}
	if len(os.Args) > 3 {
		log.Fatal("Too many arguments.")
	}

	num1, err1 := strconv.ParseInt(os.Args[1], 10, 64)
	num2, err2 := strconv.ParseInt(os.Args[2], 10, 64)
	if err1 == nil && err2 == nil {
		sum, err := adder.AddTwoInt64Numbers(num1, num2)
		if err == nil {
			fmt.Printf("%v + %v = %v \n", num1, num2, sum)
			os.Exit(0)
		} else {
			log.Fatal(err)
		}
	}
	float1, err1 := strconv.ParseFloat(os.Args[1], 4)
	float2, err2 := strconv.ParseFloat(os.Args[2], 4)
	if err1 == nil && err2 == nil {
		floatSum, err := adder.AddTwoFloatNumbers(float1, float2)
		if err == nil {
			fmt.Printf("%v + %v = %v \n", float1, float2, floatSum)
			os.Exit(0)
		} else {
			log.Fatal(err)
		}
	}
	log.Fatal("Can't parse arguments")
}
