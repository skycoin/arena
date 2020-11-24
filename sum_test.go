package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var correctValues = []struct {
	in  []string
	out float64
}{
	{
		in:  []string{"1", "2", "4"},
		out: 7.0,
	},
	{
		in:  []string{"1", "2", "4", "0", "2", "1"},
		out: 10.0,
	},
	{
		in:  []string{"1", "2", "-4", "0", "2", "1"},
		out: 2.0,
	},
	{
		in:  []string{},
		out: 0.0,
	},
	{
		in:  []string{"1000", "2000", "40", "01", "20", "1000"},
		out: 4061,
	},
	{
		in:  []string{"1000000000000.00077788", "2000", "40", "01", "20", "1000"},
		out: 1000000003061.00077788,
	},
}

var inCorrectValues = []struct {
	in []string
}{
	{
		in: []string{"o", "2", "4"},
	},
	{
		in: []string{"1", "2", "-", "0", "2", "1"},
	},
	{
		in: []string{"*"},
	},
	{
		in: []string{"@", "2000", "40", "01", "20", "1000"},
	},
}

func TestSumWithCorrectValues(t *testing.T) {
	for _, val := range correctValues {
		sum, err := Sum(val.in...)
		if err != nil {
			assert.Fail(t, "should not return error when correct values are passed")
		}
		assert.Equal(t, val.out, sum)
	}
}

func TestSumWithOneIncorrectValue(t *testing.T) {
	for _, val := range inCorrectValues {
		sum, err := Sum(val.in...)
		if err == nil {
			assert.Fail(t, "should return error when incorrect values are passed")
		}
		assert.Equal(t, 0.0, sum)
	}
}
