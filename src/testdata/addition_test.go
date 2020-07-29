package testdata

import "testing"
import "arena/src/maths"

func TestSum(t *testing.T) {
	actual := maths.Addition{FirstValue: 3, SecondValue: 4}.GetSum()
	if 7.0000 != actual {
		t.Errorf("Mistmatched data: actual: %f, expected: %f", actual, 7.000000)
	}
}
