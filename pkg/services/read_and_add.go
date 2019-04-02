package services

import (
	"fmt"
	"io"
)

// ReadNumbers reads n amount of number from stdin and returns a slice of floats
func ReadNumbers(reader io.Reader, numToRead int) ([]float64, error) {

	numbers := make([]float64, numToRead)

	for i := 0; i < numToRead; i++ {
		_, err := fmt.Fscan(reader, &numbers[i])

		if err != nil {
			return nil, err
		}

	}

	return numbers, nil

}

// AddNumbers takes slice of floats and sums them
func AddNumbers(numbers []float64) float64 {

	var sum float64

	for _, number := range numbers {
		sum += number
	}

	return sum
}
