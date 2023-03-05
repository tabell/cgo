package ex2

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Xor(a []byte, b []byte) []byte {
	j := len(a)
	k := len(b)
	var out = make([]byte, max(j, k))
	for i, _ := range a {
		out[i] = a[i%j] ^ b[i%k]
	}

	return out
}

