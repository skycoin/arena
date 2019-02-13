package main
import (
	"fmt"
	"flag"
	"strconv"
)

func main() {
	flag.Parse()
    firstNumberInt, err := strconv.Atoi(flag.Arg(0))
    if err != nil {
    	panic(err)
    }
    
    secondNumberInt, err := strconv.Atoi(flag.Arg(1))
    if err != nil {
    	panic(err)
    }

    fmt.Println("Result:", firstNumberInt + secondNumberInt)
}
