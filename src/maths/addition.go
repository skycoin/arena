package maths

type Addition struct {
	FirstValue  float64
	SecondValue float64
}

func (add Addition) GetSum() float64 {
	return add.FirstValue + add.SecondValue
}
