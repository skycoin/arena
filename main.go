package main

import (
	"fmt"
	"sumprogram/pkg/sum"
)

func main() {
	var value1, value2 int
	fmt.Scan(&value1, &value2)
	fmt.Print(sum.Sum(value1, value2))
}
