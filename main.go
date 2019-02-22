// This program takes two numbers from the command line and prints their sum.
// Currently it accepts only integer numbers.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// This is naive implementation without error handling and reflection.
func main() {
	res := 0
	nums := []string{}

	scanner := bufio.NewScanner(os.Stdin)
	for i := 1; i < 3; i++ {
		fmt.Printf("Please, enter %s number: ", num2word(i))

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				log.Println(err)
			}
			return
		}

		s := strings.TrimSpace(scanner.Text())
		if len(s) == 0 {
			log.Fatalln("Empty value received. Exiting")
		}
		nums = append(nums, s)

		v, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalln(err)
		}

		res += v
	}

	fmt.Printf("The sum of %s and %s is: %d\n", nums[0], nums[1], res)
}

// num2word converts numeric representation of a counter to verbal.
func num2word(num int) string {
	switch num {
	case 1:
		return "first"
	case 2:
		return "second"
	default:
		return "the next"
	}
}
