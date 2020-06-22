package main

import (
	"bufio"
	"fmt"
	"io"
)

// MakeOperation fnction reads two numbers from console, add them and return result
func MakeOperation(rd io.Reader) (num1, num2, res float64, err error) {
	reader := bufio.NewReader(rd)
	fmt.Print("\n Enter first number by space:")
	_, err = fmt.Fscan(reader, &num1, &num2)
	if err != nil {
		return 0, 0, 0, err
	}
	res = num1 + num2
	return num1, num2, res, nil
}
