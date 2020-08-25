package main
 
import "fmt"
 
func main() {
   fmt.Print("Enter First Number: ")   
   var first int
   fmt.Scanln(&first)                  // Take input from user
   fmt.Print("Enter Second Number: ")
   var second int
   fmt.Scanln(&second)
   fmt.Print(first + second)           // Addition of two numbers   
}
