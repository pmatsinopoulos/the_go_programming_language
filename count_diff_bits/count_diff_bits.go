package count_diff_bits

func CountDifferentBits(a [32]byte, b [32]byte) int {
	var masks [8]byte = [8]byte{
		0x01,
		0x02,
		0x04,
		0x08,
		0x10,
		0x20,
		0x40,
		0x80,
	}

	var diff = 0
	for i := range a {
		for j := range 8 {
			var aa = a[i] & masks[j]
			var bb = b[i] & masks[j]
			if aa != bb {
				diff += 1
			}
		}
	}

	return diff
}
