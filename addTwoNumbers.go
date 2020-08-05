//Run  go run addTwoNumbers.go --num1 12 --num2 24

package main

import (
	"flag"
	"fmt"
	"os"
)
func addTwoNumbers(num1 int64,num2 int64)int64{
	return num1+num2
}

func main() {

	num1 := flag.Int64("num1", 0, "First Number ")

	num2 := flag.Int64("num2", 0, "Second Number")


	flag.Parse()
	if *num1==0 || *num2==0{
		fmt.Println("Numbered Entered shouldnt be 0")
		os.Exit(1)
	}
	fmt.Println("First Number:", *num1)
	fmt.Println("Second Number:", *num2)
	fmt.Println("Sum of two numbers is ", addTwoNumbers(*num1,*num2))


}