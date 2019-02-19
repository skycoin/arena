package main

import (
	"arena/pkg"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	Error := pkg.AddendWrong(args)

	if Error.Val > 0 {
		switch Error.Val {
		case 1:
			fmt.Println("Addends does not indicated")

		case 2:
			fmt.Println("Count Addends must be 2")

		case 3:
			fmt.Println("Addends indicated more than 2")
		}
		return
	}


	addn1, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("First arg is not integer. %v\n", err)
		return
	}

	addn2, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("Second arg is not integer. %v\n", err)
		return
	}

	var a = pkg.Addend{
		AddnA: addn1,
		AddnB: addn2,
	}
	a.Adds()

	fmt.Printf("Summ Result %v\n", a.AddnS)
}
