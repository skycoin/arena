package main
import (
	"fmt"
	"flag"
	"strconv"
)

package main
import (
	"fmt"
	"flag"
	"strconv"
)

func main() {
	add()
}

func add() error {
    flag.Parse()
    firstNumberInt, err := strconv.Atoi(flag.Arg(0))
    if err != nil {
    	return err
    }

    secondNumberInt, err := strconv.Atoi(flag.Arg(1))
    if err != nil {
    	return err
    }

    fmt.Println("Result:", firstNumberInt + secondNumberInt)
    return nil
}
