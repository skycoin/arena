package main

import (
	"fmt"
	"github.com/orochyy/sumprogram/pkg/sum"
)

func main() {
	var v1, v2 int
	fmt.Printf("v1: ")
	fmt.Scan(&v1)
	fmt.Print(" \nv2: ")
	fmt.Scan(&v2)
	fmt.Print(sum.Sum(v1, v2))
}
