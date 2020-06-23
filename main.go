package main

import (
	"fmt"
)

func main() {
	var f float64
	fmt.Println("Enter first number : ")
	_, err := fmt.Scanf("%f", &f)
	if err != nil {
		fmt.Println(err)
	}
	var n float64
	fmt.Println("Enter second number : ")
	_, err = fmt.Scanf("%f", &n)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("You have entered : %f \n", f+n)

}
