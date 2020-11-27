package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	addToNumbers()
}

func addToNumbers() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number then press enter: ")
	str1, err := reader.ReadString('\n')
	if err != nil {
		log.Println("got error:", err)
		os.Exit(1)
	}
	num1, err := strconv.Atoi(strings.TrimSpace(str1))
	if err != nil {
		log.Printf("cannot conver %s to int, got error %v", str1, err)
		os.Exit(1)
	}

	fmt.Print("Enter another number then press enter: ")
	str2, err := reader.ReadString('\n')
	if err != nil {
		log.Println("got error:", err)
		os.Exit(1)
	}
	num2, err := strconv.Atoi(strings.TrimSpace(str2))
	if err != nil {
		log.Printf("cannot conver %s to int, got error %v", str2, err)
		os.Exit(1)
	}

	sum := num1 + num2
	fmt.Printf("Result: %d + %d = %d\n", num1, num2, sum)
}
