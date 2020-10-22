package main

import (
	"fmt"
	"strconv"
)

const ERROR  = "Please Enter Integer Value Only"

func main()  {
	fmt.Println("Enter First Number: ")
	var fstr string
	fmt.Scanln(&fstr)
	first,err := strconv.Atoi(fstr)
	if err  != nil{
		fmt.Println(ERROR)
		return
	}

	fmt.Print("Enter Second Number: ")
	var lstr string
	fmt.Scanln(&lstr)
	second,err := strconv.Atoi(lstr)
	if err  != nil{
		fmt.Println(ERROR)
		return
	}

	sum := first + second
	fmt.Println("Sum :",sum)
}
