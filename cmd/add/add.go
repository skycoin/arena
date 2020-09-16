package main

import (
	"fmt"
	"log"
	"os"

	"github.com/skycoin/arena/pkg/math"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal(math.NumError)
	}
	first, second := os.Args[1], os.Args[2]
	res, err := math.AddNumbers(first, second)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(os.Stdout, "The result of adding", first, "and", second, "is", res)
}
