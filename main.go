package main

import (
	"fmt"
	"log"
	"skycoin/arena/pkg"
)

func main() {
	sum, err := pkg.AddTwoNums()
	if err != nil {
		log.Fatalf("Error adding sum %v", err)
	}
	fmt.Printf("The sum is %v\n", sum)
}
