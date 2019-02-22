package arena

import (
	"fmt"
	"strconv"
)

// IntOp - operation with int
type IntOp interface {
	ExecInt() (int, error)
}

// ExecInt - execute int operation
func ExecInt(o IntOp) (int, error) {
	return o.ExecInt()
}

// Add2 - sum two int
func Add2(o ...IntOp) IntOp {
	if len(o) != 2 {
		return Error{err: fmt.Errorf("wrong count of args: not equals 2 (actual count: %d)", len(o))}
	}
	return sum(o[:2])
}

type sum []IntOp

// ExecInt - execute int operation
func (s sum) ExecInt() (int, error) {
	var out int
	for i, o := range s {
		v, err := o.ExecInt()
		if err != nil {
			return 0, fmt.Errorf("error on %d arg: %v", i+1, err)
		}
		out += v
	}
	return out, nil
}

// Int number
type Int int

// ExecInt - execute int operation
func (i Int) ExecInt() (int, error) {
	return int(i), nil
}

// StrInt - int in string representation
type StrInt string

// ExecInt - execute int operation
func (s StrInt) ExecInt() (int, error) {
	i, err := strconv.Atoi(string(s))
	if err != nil {
		return 0, fmt.Errorf("%q is not a number: %v", string(s), err)
	}
	return i, nil
}

// Error of operation
type Error struct {
	err error
}

// ExecInt - execute int operation
func (e Error) ExecInt() (int, error) {
	return 0, e.err
}
