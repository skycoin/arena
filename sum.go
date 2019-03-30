package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var num1, num2 int

	fmt.Println("enter a number:")

	if _, err := fmt.Scan(&num1); err != nil {
		log.Print("  Scan for num failed, due to ", err)
		return
	}

	fmt.Println("enter a second number:")

	if _, err := fmt.Scan(&num2); err != nil {
		log.Print("  Scan for num2 failed, due to ", err)
		return
	}

	result := fmt.Sprintf("%d + %d = %d", num1, num2, num1+num2)
	io.WriteString(os.Stdout, result) //
}