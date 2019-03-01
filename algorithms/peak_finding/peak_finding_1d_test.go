package algorithms

import (
	"testing"
)

func TestPeakFind1D(t *testing.T) {
	cases := []struct {
		array  []int
		result int
	}{
		{[]int{1, 2, 3, 4, 3, 5},
			4,
		},
		{
			[]int{2},
			2,
		},
		{
			[]int{1, 4, 10, 4, 1, 5},
			10,
		},
		{
			[]int{0, 1, 0, 2, 0},
			1,
		},
	}

	for _, testCase := range cases {
		index := peakFind(testCase.array, len(testCase.array))
		result := testCase.array[index]
		if result != testCase.result {
			t.Error("Peak found is different than ", testCase.array, index, result)
		}
	}
}
