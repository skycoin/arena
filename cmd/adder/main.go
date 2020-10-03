package main

import (
	"fmt"
	"log"
	"os"

	"github.com/skycoin/arena/pkg/adder"
)

func main() {
	a, b, err := adder.ParseArguments(os.Args[1:])
	if err != nil {
		log.Fatal(err.Error())
	}
	if a.IsInf() || b.IsInf() {
		log.Fatal("Error: infinities aren't allowed")
	}
	res := adder.AddNumbers(a, b)
	fmt.Printf("Result: %s\n", res.String())
}
