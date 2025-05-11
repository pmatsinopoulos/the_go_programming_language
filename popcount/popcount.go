package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (i.e. number of set bits) of x.
// A uint64 is 64 bits, i.e. 8 bytes long.
func PopCount(x uint64) int {
	// each pc[value] where 0 <= value <= 255 gives me the number
	// of bits set for that value.
	// Hence for the 8 bytes of the uint64 I just add the number
	// of bits returned by pc for each one of the bytes.
	//
	return int(
		pc[byte(x>>(0*8))] + pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] + pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] + pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] + pc[byte(x>>(7*8))],
	)
}
