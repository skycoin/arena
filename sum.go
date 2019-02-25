package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter first numbers:")
	finput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	a, err := strconv.Atoi(strings.TrimSpace(finput))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Enter second numbers:")
	sinput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := strconv.Atoi(strings.TrimSpace(sinput))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Sum is :", a+b)
}
