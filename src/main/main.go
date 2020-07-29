package main

import (
	"awesomeProject1/src/maths"
	"fmt"
)

func main() {
	add := new(maths.Addition)
	var value1 float32
	var value2 float32
	fmt.Println("Enter first number")
	fmt.Scanf("%g", &value1)
	fmt.Println("Enter Second number")
	fmt.Scanf("%g", &value2)
	add.FirstValue = value1
	add.SecondValue = value2
	fmt.Printf("Addition of two no's %g:", add.GetSum())
}
