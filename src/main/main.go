package main

import (
	"awesomeProject1/src/maths"
	"fmt"
)

func main() {
	add := new(maths.Addition)
	var value1 float32
	var value2 float32
	fmt.Println("Enter two numbers")
	_, err := fmt.Scanf("%g, %g", &value1, &value2)
	if err != nil {
		add.FirstValue = value1
		add.SecondValue = value2
		fmt.Printf("Addition of two no's %g:", add.GetSum())
	} else {
		fmt.Errorf("Error:", err)
	}
}
