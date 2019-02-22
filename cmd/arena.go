package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Add2Number takes two numbers from the command line, adds them up
// If one of the number contains non-digit letter, print 0
func Add2Number(reader io.Reader) {
	var number1, number2 float64
	fmt.Fscanf(reader, "%f %f", &number1, &number2)
	fmt.Println(number1 + number2)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	Add2Number(reader)
}
