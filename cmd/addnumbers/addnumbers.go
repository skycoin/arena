package main

import (
	"arena/calc"
	"arena/input"
	"fmt"
	"math/big"
)

func main() {
	var firstBigNumber, secondBigNumber *big.Float
	fmt.Println("Enter Two numbers two be Added : ")
	firstBigNumber = input.Input("Enter First Number: ")
	secondBigNumber = input.Input("Enter Second Number: ")
	fmt.Println("Adding ....")
	var addstruct = calc.AddStruct{
		X: firstBigNumber,
		Y: secondBigNumber,
	}
	result := calc.Add(&addstruct)
	fmt.Println("Result", result)
}
