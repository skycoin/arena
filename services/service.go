package services

type ServiceInterface interface {
	SumsUp() (int, error)
}

type serviceInput struct {
	input []int
}

// New ..
func New(input []int) ServiceInterface {

	return &serviceInput{
		input: input,
	}
}

func (s serviceInput) SumsUp() (int, error) {
	i := 0
	for _, k := range s.input {
		i = i + k
	}
	return i, nil
}
