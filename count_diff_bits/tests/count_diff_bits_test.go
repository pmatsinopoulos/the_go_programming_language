package tests

import "testing"
import "the_go_programming_language/count_diff_bits"

func TestCountDifferentBits(t *testing.T) {
	testCases := []struct {
		name     string
		a        [32]byte
		b        [32]byte
		expected int
	}{
		{
			name:     "Identical inputs",
			a:        [32]byte{0x01, 0x02},
			b:        [32]byte{0x01, 0x02},
			expected: 0,
		},
		{
			name:     "Two bits difference",
			a:        [32]byte{0x01, 0x02},
			b:        [32]byte{0x01, 0x01},
			expected: 2,
		},
		{
			name:     "Five bits difference",
			a:        [32]byte{0x01, 0x02, 0x11},
			b:        [32]byte{0x01, 0x01, 0x38},
			expected: 5,
		},
		{
			name:     "11 bits difference",
			a:        [32]byte{0x01, 0x02, 0x11, 0xFa, 0xeb},
			b:        [32]byte{0x01, 0x01, 0x38, 0xFa, 0x1e},
			expected: 11,
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name, func(t *testing.T) {
				result := count_diff_bits.CountDifferentBits(tc.a, tc.b)
				if result != tc.expected {
					t.Errorf("%s: got %d, want %d", tc.name, result, tc.expected)
				}
			},
		)
	}
}
