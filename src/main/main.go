package main

import (
	"arena/src/maths"
	"fmt"
	"os"
	"strconv"
)

func main() {
	add := new(maths.Addition)
	add.FirstValue, _ = strconv.ParseFloat(os.Args[0], 64)
	add.SecondValue, _ = strconv.ParseFloat(os.Args[1], 64)
	fmt.Printf("Addition of two no's %g:", add.GetSum())
}
