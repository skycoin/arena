package testdata

import (
	"arena/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

var flagTests = []struct {
	addnA int
	addnB int
	addnS int
}{
	{1, 01, 2},
	{0, 1, 1},
	{1, 0, 1},
}

func TestAdds(t *testing.T) {
	var addn pkg.Addend

	for _, addns := range flagTests {
		addn.AddnA = addns.addnA
		addn.AddnB = addns.addnB
		addn.Adds()
		assert.Equal(t, addn.AddnS, addns.addnS, "Adds method is incorrect")
	}
}
