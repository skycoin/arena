package arena

import (
	"fmt"
	"io"
)

// ReadNumbers reads number inputs
func ReadNumbers(reader io.Reader) (float64, float64, error) {
	numbers := [2]float64{}
	for i := 0; i < 2; i++ {
		_, err := fmt.Fscan(reader, &numbers[i])
		if err != nil {
			return 0.0, 0.0, err
		}
	}
	return numbers[0], numbers[1], nil
}

// AddNumbers adds two numbers
func AddNumbers(a, b float64) float64 {
	sum := a + b
	return sum
}
