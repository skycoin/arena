package main

import (
	"arena/pkg"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	addn1, err := strconv.Atoi(args[0])
	addn2, err := strconv.Atoi(args[1])

	if err != nil {

	}
	var a = pkg.Addend{
		AddnA: addn1,
		AddnB: addn2,
	}
	a.Adds()

	fmt.Printf("Summ Result %v\n", a.AddnS)
}
