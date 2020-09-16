package main

import (
	"fmt"
)

func Addition() {
	// fmt.Println("function part")

	var First_number, Second_number int

	fmt.Println("Enter your First Number:")
	fmt.Scanln(&First_number)

	fmt.Println("Enter your Second Number:")
	fmt.Scanln(&Second_number)

	Result := First_number + Second_number

	fmt.Println("Addition of two numbers:", First_number, "+", Second_number, "=", Result)
	fmt.Println("Result :", Result)

}

func main() {
	// fmt.Println("main func")
	Addition()
}
