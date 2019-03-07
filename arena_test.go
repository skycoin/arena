package main

/*
 Excerpts from: https://github.com/skycoin/skycoin/wiki/Idiomatic-Go-(in-Skycoin)#testing

 - Use the `testify require` or `testify assert` package for assertions.
 - Write table-driven tests when appropriate; most tests can be written this way.
*/
import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

/*
Although Idiomatic Go (in Skycoin) stated:
  - Use struct field names when declaring a struct object instance
  - Use the struct's field names when instantiating it, and put them each on their own line

 I decided to go with format described in recommended
wiki article https://github.com/golang/go/wiki/TableDrivenTests

 I hope that this deviation from rules is acceptable for table-driven tests
*/
var cases = []struct{ argx, argy, result string }{
	{"1", "1", "2"},
	{"-1", "1", "0"},
	{"10e6", "-Inf", "-Inf"},
	{"10e-10", "2.2e-10", "1.2200000000000001e-09"},
	{"10e120", "0", "1e+121"},
	{"Inf", "Inf", "+Inf"},
	{"Inf", "-Inf", "NaN"},
	{"A", "10", "strconv.ParseFloat: parsing \"A\": invalid syntax"},
	{"10", "B", "strconv.ParseFloat: parsing \"B\": invalid syntax"},
}

func TestArena(t *testing.T) {
	for _, tt := range cases {
		testname := fmt.Sprintf("%v + %v = %v", tt.argx, tt.argy, tt.result)
		t.Run(testname, func(t *testing.T) {
			require.Equal(t, sumargs(tt.argx, tt.argy), tt.result)
		})
	}
}
