// Golang program to show how
// to take input from the user and produce a sum
package main

import (
	"fmt"
	"log"
)

const (
	msz1 = "Enter Your First Number: "
	msz2 = "Enter Your Second Number: "
)

type Inputs struct {
	A, B float64
}

func (i *Inputs) Sum() float64 {
	return i.A + i.B
}

func main() {

	inputs := new(Inputs)
	fmt.Print(msz1)
	if _, err := fmt.Scanln(&inputs.A); err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Print(msz2)
	if _, err := fmt.Scanln(&inputs.B); err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Printf("Your Result is: %v\n", inputs.Sum())
}
