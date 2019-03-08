package main

import (
	"flag"
	"fmt"

	"github.com/arena/cmd"
)

func main() {
	// Parse operation
	var operation int
	flag.IntVar(&operation, "operation", 1, "Invalid operation !!!")
	flag.Parse()
	cli := cmd.Cmd{}

	// Can add more operations.
	switch operation {
	case 1:
		cli.Add()
	default:
		fmt.Println("Invalid operation !!!")
	}
}
