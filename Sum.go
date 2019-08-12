package main
import "fmt"
import "os"
import "strconv" 
func main(){
arg1 := os.Args[1]
arg2 := os.Args[2]

n1,err1 := strconv.Atoi(arg1)
if err1 == nil {
fmt.Println(n1)
}

n2,err2 := strconv.Atoi(arg2)
if err2 == nil {
fmt.Println(n2)
}

sum := n1+n2;
fmt.Println("sum : ",sum);
}
