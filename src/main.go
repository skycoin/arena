package main

import (
	"fmt"
	"os"

	sum "github.com/skycoin/arena/pkg"
)

func main() {
	args := os.Args[1:]
	res := sum.ArgAddition(args)

	fmt.Printf("%d\n", res)

}
