package main

import (
	"fmt"
	"log"
	"os"

	"github.com/skycoin/arena/src/additionTask"
)

func main() {
	//Reading command line arguments
	args := os.Args[1:]
	sum, err := additionTask.AdditionOfTwoNums(args)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Sum of %s and %s is %f \n", args[0], args[1], sum)
}
