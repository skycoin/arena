package main

import (
	"fmt"
	"flag"
)

func main(){
	firstArg := flag.Int("num1", 0, "First Input")
	secondArg := flag.Int("num2", 0, "Second Input")

	flag.Parse()

	result := *firstArg + *secondArg
	fmt.Println(result)
}
