package main

import (
	"fmt"
	"os"
	"strconv"
)

// Galymzhan Issabekov
//Telegram: isabekovgg
//email: isabekovgg@gmail.com
//code writen in 5 minutes
func main() {
	arg := os.Args[1:]
	if len(arg) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	if len(arg) < 2 {
		fmt.Println("Not enough arguments")
		return
	}
	//it was no rules, to not use libraries, that's why i used atoi, if it's needed i can provide my own atoi function
	n1, err1 := strconv.Atoi(arg[0])
	n2, err2 := strconv.Atoi(arg[1])
	if err1 == nil && err2 == nil {
		res := n1 + n2
		fmt.Println(res)
	}
}
