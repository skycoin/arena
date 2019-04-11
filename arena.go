package arena

import "fmt"

// PrintSum adds numbers and prints their result
func PrintSum(nums ...float64) {
	sum := float64(0)
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
}
