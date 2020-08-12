package main

import (
	"fmt"
	"os"

	"github.com/skycoin/arena/src/parser"
)

func main() {
	args := os.Args[1:]
	argParser := parser.NewArgumentParser()
	num1, num2, err := argParser.Take2Numbers(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num1 + num2)
}
