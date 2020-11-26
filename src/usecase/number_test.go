// +build unit

package usecase

import (
	"math/rand"
	"testing"

	"github.com/bobbiprathama/arena/src/helpers"
	"github.com/stretchr/testify/require"
)

func TestNumber_Sum_ShouldReturnNilParamError_WhenGivenNilNumberParam(t *testing.T) {

	// -- init
	fakeNumber1 := helpers.IntToInt64(nil)
	fakeNumber2 := helpers.IntToInt64(helpers.IntPointer(rand.Int()))

	// -- CODE UNDER TEST
	uc := NewNumberOperation()
	actualRes, err := uc.Sum(fakeNumber1, fakeNumber2)

	// -- expectations
	require.Nil(t, actualRes)

	// -- should return error
	require.Error(t, err)
	expectedError := NewNilParamError()
	require.Equal(t, err, expectedError)
}

func TestNumber_Sum_ShouldReturnSumResult_WhenGivenNonNilNumberParams(t *testing.T) {

	var expectations = []struct {
		input1 int64
		input2 int64
		output int64
	}{
		{
			int64(10),
			int64(12),
			int64(22),
		},
		{
			int64(2),
			int64(2),
			int64(4),
		},
		{
			int64(0),
			int64(5),
			int64(5),
		},
	}

	uc := NewNumberOperation()
	for _, exp := range expectations {
		actual, err := uc.Sum(&exp.input1, &exp.input2)
		require.NoError(t, err)
		require.NotNil(t, actual)
		require.Equal(t, exp.output, *actual)
	}

}
