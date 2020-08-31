package main

import "fmt"
import "strings"
import "strconv"

var num1 string
var num2 string

func add(res1 float64, res2 float64) {

	fmt.Println("%%%%%%%%%%% Result Report %%%%%%%%%%%%%%")
	fmt.Printf("Your first number: %v\n", res1)
	fmt.Printf("Your second number: %v\n", res2)

	finalres := res1 + res2
	fmt.Printf("Result: %v\n", finalres)
	fmt.Println("=========== Tank you, Bye! ==============")
}

func main() {
	fmt.Println("=========== Add 2 numbers program ==============")
	fmt.Println("Please enter your numbers to be added...")

	fmt.Print("Enter first number please: ")
	fmt.Scanf("%s", &num1)
	num1Trim := strings.TrimSpace(num1)
	res1, _ := strconv.ParseFloat(num1Trim, 8)

	fmt.Print("Enter second number please: ")
	fmt.Scanf("%s", &num2)
	num2Trim := strings.TrimSpace(num2)
	res2, _ := strconv.ParseFloat(num2Trim, 8)

	add(res1, res2)
}
