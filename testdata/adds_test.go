package testdata

import (
	"testing"
	"arena/pkg"
	"github.com/stretchr/testify/assert"
)

var flagtests = []struct {
	addnA  int
	addnB  int
	addnS  int
}{
	{1, 01, 2},
	{0, 1, 1},
	{1,0, 1},
}

func TestFlagParser(t *testing.T) {
	var addn pkg.Addend

	for _, addns := range flagtests {
		addn.AddnA = addns.addnA
		addn.AddnB = addns.addnB
		addn.Adds()
		assert.Equal(t, addn.AddnS, addns.addnS, "Adds method is incorrect")
	}
}
