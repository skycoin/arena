package main
import(
  "fmt"
"strconv"
)

func main(){
  fmt.Println("Please enter 2 int numbers: ")
  var num1,num2 string
  fmt.Scanln(&num1,&num2)
  n1,_:=strconv.Atoi(num1)
  n2,_:=strconv.Atoi(num2)
  sum:=strconv.Itoa(n1+n2)
  fmt.Printf("the sum of these two numbers is: %v\n",sum)
}
