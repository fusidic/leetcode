package netease

func findKthBit(n int, k int) byte {
	S := []byte{'a'}
	for i := 2; i <= n; i++ {
		nextS := append(S, byte(i-1+'a'))
		nextS = append(nextS, reverse(invert(S))...)
		S = nextS
	}
	return S[k-1]
}

func reverse(s []byte) []byte {
	l := len(s)
	if l == 1 {
		return s
	}
	for i := 0; i < l/2; i++ {
		s[i], s[l-1-i] = s[l-1-i], s[i]
	}
	return s
}

func invert(s []byte) []byte {
	l := len(s)
	for i := 0; i < l; i++ {
		s[i] = helper(s[i])
	}
	return s
}

func helper(a byte) byte {
	x := int(a - 'a')
	return byte(25 - x + 'a')
}
