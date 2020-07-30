package maths

type Addition struct {
	FirstValue  float32
	SecondValue float32
}

func (add Addition) GetSum() float32 {
	return add.FirstValue + add.SecondValue
}
