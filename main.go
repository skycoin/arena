package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const HowManyInputNumber int = 2

// SumUpNumber is function to sum up number
// num is array of int
func SumUpNumber(num []int) int {
	var r int
	for _, val := range num {
		r += val
	}
	return r
}

// GetInputNumber is function to get input number form user
// reader is buffer input
func GetInputNumber(reader *bufio.Reader) int {
	var (
		userText string
		num      int
		err      error
	)
	if userText, err = reader.ReadString('\n'); err != nil {
		fmt.Print("Please Enter Only Number: ")
		return GetInputNumber(reader)
	}
	// trim \n from input
	if num, err = strconv.Atoi(strings.TrimSpace(userText)); err != nil {
		fmt.Print("Please Enter Only Number: ")
		return GetInputNumber(reader)
	}
	return num
}

func main() {
	taskTest := fmt.Sprintf(`
---------------------------------------------------------------
Write a Golang program that takes two numbers from the command line, adds them up and prints the result.
---------------------------------------------------------------`)
	fmt.Println(taskTest)

	var (
		reader = bufio.NewReader(os.Stdin)
		t      []int
	)
	for i := 1; i < HowManyInputNumber+1; i++ {
		fmt.Print(`Input Number: `)
		n := GetInputNumber(reader)
		t = append(t, n)

	}
	result := fmt.Sprintf(`Result for %s is %d`,
		strings.Trim(strings.Replace(fmt.Sprint(t), " ", " + ", -1), "[]"),
		SumUpNumber(t))
	fmt.Println(result)
}
