package reverse_ints

func Reverse(s []int) {
	for i := range len(s) / 2 {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}
