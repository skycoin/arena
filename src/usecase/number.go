package usecase

type NumberOperation struct {
}

func NewNumberOperation() NumberOperation {
	return NumberOperation{}
}

func (op NumberOperation) Sum(x, y *int64) (*int64, error) {
	if x == nil || y == nil {
		return nil, NewNilParamError()
	}

	sumRes := *x + *y
	return &sumRes, nil
}
