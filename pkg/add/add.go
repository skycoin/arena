package add

// Data is to scan two numbers for addition
type Data struct {
	Numbers []float64
}

// Add function will retur the sum of two numbers
func (d *Data) Add() float64 {
	var sum float64
	for _, number := range d.Numbers {
		sum += number
	}
	return sum
}
