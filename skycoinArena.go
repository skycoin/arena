package main

import (
    "fmt"
)

func main (){
	fmt.Print("Enter first number:")
    var first float32
	fmt.Scanln(&first)
	fmt.Print("Enter second number:")
    var second float32
    fmt.Scanln(&second)
    fmt.Print(first + second)
}