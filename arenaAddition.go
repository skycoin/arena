package main

import (
    	"fmt"
   	 )

func main() {

    	//addition func call	
    	Addition()
}


func Addition() {

	fmt.Println("\n")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("						 ")
	fmt.Println("	Addition function!!		         ")
	fmt.Println("						 ")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	fmt.Println("\n")

    	fmt.Println("Enter Below The First Number:")
    	var numOne int
    	fmt.Scanln(&numOne)
	
    	fmt.Println("Enter Below The Second Number:")
    	var numTwo int
    	fmt.Scanln(&numTwo)
	fmt.Println("\n")

    	fmt.Println("The Addition of Both Numbers is = ", numOne + numTwo)    
	fmt.Println("\n")

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("						 ")
	fmt.Println("	Good Bye!!		         	 ")
	fmt.Println("						 ")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("\n")

}