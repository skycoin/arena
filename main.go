package main

import (
	"fmt"
	"log"
	"os"

	"github.com/skycoin/arena/pkg/skytask"
)

func main() {
	sum, err := skytask.Add2Num(os.Args...)
	if nil != err {
		log.Fatal(err)
	}
	fmt.Printf("You got a sum of %f\n", sum)
}
