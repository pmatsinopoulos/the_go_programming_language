package append_int

import (
	"testing"
)

func TestAppendIt(t *testing.T) {
	testCases := []struct {
		name          string
		inputSlice    []int
		inputInt      int
		expectedSlice []int
	}{
		{
			name:          "First case",
			inputSlice:    make([]int, 3, 4),
			inputInt:      5,
			expectedSlice: []int{0, 0, 0, 5},
		},
		{
			name:          "Second case",
			inputSlice:    []int{1, 3, 8},
			inputInt:      5,
			expectedSlice: []int{1, 3, 8, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := AppendInt(tc.inputSlice, tc.inputInt)

			if !equalSlices(result, tc.expectedSlice) {
				t.Errorf("%s, Expected %v, got %v", tc.name, tc.expectedSlice, result)
			}
		})
	}
}

func equalSlices(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}
	return true
}
