package reverse_ints

import (
	"slices"
	"testing"
)

func TestReverseInts(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Reverse 3",
			input:    []int{28, 30, 13},
			expected: []int{13, 30, 28},
		},
		{
			name:     "Reverse 5",
			input:    []int{28, 30, 13, 20, 2},
			expected: []int{2, 20, 13, 30, 28},
		},
		{
			name:     "Reverse 1",
			input:    []int{28},
			expected: []int{28},
		},
		{
			name:     "Reverse none",
			input:    []int{},
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// we preserve the original input
			inputCopy := slices.Clone(tc.input)

			Reverse(inputCopy)

			if !equalSlices(inputCopy, tc.expected) {
				t.Errorf("Expected %v, result was %v\n", tc.expected, tc.input)
			}
		})
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
