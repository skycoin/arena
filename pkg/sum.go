package sum

import (
	"strconv"
)

// ArgAddition sums two Command line arguments
func ArgAddition(args []string) int {
	sum := 0
	for _, a := range args {
		num, _ := strconv.Atoi(a)
		sum += num
	}

	return sum
}
