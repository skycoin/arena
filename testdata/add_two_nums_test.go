package main

import (
	"kycoin/arena/pkg"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestAddTwoNums(t *testing.T) {

	sum, err := pkg.AddTwoNums()
	if err != nil {
		t.Errorf("Test failed %v", err)
	}
	assert.Equal(t, sum, 10, "The sum should be 10")
}
