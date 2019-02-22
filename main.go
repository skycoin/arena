package main

import (
    "bufio"
    "fmt"
	"os"
	"strings"
	"strconv"
	"runtime"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	//get the first input from user
	fmt.Print("Enter a number: ")
	num1, _ := reader.ReadString('\n')

	//replace unexpectd character to avoid error while convert to int 
	if runtime.GOOS == "windows" {
		num1 = strings.Replace(num1, "\r\n", "", -1)
	}else{
		num1 = strings.Replace(num1, "\n", "", -1)
	}
	
	//convert user's input to int
	input1, err := strconv.Atoi(num1)
    if err != nil {
        fmt.Println("Plese enter only number, Try again")
		os.Exit(2)
	}

	fmt.Print("Enter another number: ")
	num2, _ := reader.ReadString('\n')

	//replace unexpectd character to avoid error while convert to int
	if runtime.GOOS == "windows" {
		num2 = strings.Replace(num2, "\r\n", "", -1)
	}else{
		num2 = strings.Replace(num2, "\n", "", -1)
	}

	//convert user's input to int
	input2, err := strconv.Atoi(num2)
    if err != nil {
        fmt.Println("Plese enter only number, Try again")
		os.Exit(2)
	}
	
	//print values
	fmt.Println("Input number 1: ", input1)
	fmt.Println("Input number 2: ", input2)
	fmt.Println("Total vaule is: ", input1 + input2)
}
