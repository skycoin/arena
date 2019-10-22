package main

import "log"
import "fmt"
import "os"
import "strconv"

func main() {

	if len(os.Args) < 3 {
		log.Fatal("Not enough arguments.")
	}
	if len(os.Args) > 3 {
		log.Fatal("Too many arguments.")
	}
	var num1, num2 int64
	num1, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	num2, err = strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	sum := num1 + num2
	fmt.Println(sum)
}