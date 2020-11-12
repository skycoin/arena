package test

import (
	"strconv"
	"testing"

	sum "github.com/skycoin/arena/pkg"
)

func TestMain(t *testing.T) {
	// call flag.Parse() here if TestMain uses flags
	t.Run("MainTest", MainTest)
}

func MainTest(t *testing.T) {

	arg1 := 2
	arg2 := 3

	args := []string{"go run ./ src",
		strconv.FormatInt(int64(arg1), 10),
		strconv.FormatInt(int64(arg2), 10),
	}

	got := sum.ArgAddition(args[1:])
	want := 5

	if got != want {
		t.Errorf("%d + %d = %d; want %d", arg1, arg2, got, want)
	}
}
