package append_int

func AppendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1

	if cap(x) >= zlen {
		z = x[:zlen]
	} else {
		// not much capacity on destination
		zcap := zlen * 2
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[zlen-1] = y

	return z
}
