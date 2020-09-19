package main

import (
	"fmt"
	"log"
	"os"

	"github.com/skycoin/arena/src/skycoinArenaTask"
)

func main() {
	sum, AddErr := skycoinArenaTask.Add(os.Args...)
	if AddErr != nil {
		log.Fatal(AddErr)
	}
	fmt.Printf("Sum: %f\n", sum)
}
