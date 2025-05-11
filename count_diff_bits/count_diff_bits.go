package count_diff_bits

import "math/bits"

func CountDifferentBits(a *[32]byte, b *[32]byte) int {
	var diff = 0
	for i := range a {
		diff += bits.OnesCount8(a[i] ^ b[i])
	}

	return diff
}
