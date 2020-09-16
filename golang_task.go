package main

import (
	"fmt"
)

//Addition function Declaration
func Addition() {

	var First_number, Second_number int // Variable declaration

	fmt.Println("Enter your First Number:")
	fmt.Scanln(&First_number)

	fmt.Println("Enter your Second Number:")
	fmt.Scanln(&Second_number)

	Result := First_number + Second_number // Results for addition

	fmt.Println("Addition of two numbers:", First_number, "+", Second_number, "=", Result)
	fmt.Println("Result :", Result)

}

// Main function declaration
func main() {

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("                                                 ")
	fmt.Println("                                                 ")
	fmt.Println("*************************************************")
	fmt.Println("                                                 ")
	fmt.Println("             Adding Two Numbers                                    ")
	fmt.Println("                                                 ")
	fmt.Println("*************************************************")
	fmt.Println("                                                 ")
	fmt.Println("                                                 ")

	Addition() //function call in main function

	fmt.Println("                                                 ")
	fmt.Println("                                                 ")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

}
