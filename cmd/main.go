package main

import (
	"fmt"
	"os"

	"github.com/skycoin/arena/src/parser"
)

func main() {
	args := os.Args[1:]
	argParser := parser.NewArgumentsParser()
	n1, n2, err := argParser.Take2Numbers(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n1 + n2)
}
