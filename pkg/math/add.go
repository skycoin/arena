package math

// Add returns the sum of input numbers
func Add(numbers ...float64) float64 {
	var result float64

	for _, number := range numbers {
		result += number
	}

	return result
}
