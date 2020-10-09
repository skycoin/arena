package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Please provide exact two arguments")
	} else {
		a,err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		b,err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("sum of %d+%d  = %d\n",a,b, a+b)
	}
}
