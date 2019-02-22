package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/skycoin/arena/pkg/sum"
)

func main() {
	var err error
	var a, b int64

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter two integer numbers to get their sum")
	fmt.Println("Enter first number:")
	scanner.Scan()
	a, err = strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println("Enter second number:")
	scanner.Scan()
	b, err = strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println("Sum =", sum.Int64(a, b))
}
