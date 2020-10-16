package main
 
import "fmt"
 
func main() {
	

	var first int
	var second int
        
	fmt.Print("Enter First Digit: ")             
        fmt.Scanln(&first)                  
    

	fmt.Print("Enter Second Digit: ")
        fmt.Scanln(&second)
        

	fmt.Println("Sum of digits",first + second)            
}