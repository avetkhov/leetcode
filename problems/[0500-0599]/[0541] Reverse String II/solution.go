package main

func reverseStr(s string, k int) string {
	b := []byte(s)

	for i := 0; i < len(b); i += 2 * k {
		l, r := i, i+k-1

		if r >= len(b) {
			r = len(b) - 1
		}

		for l < r {
			b[l], b[r] = b[r], b[l]
			l++
			r--
		}
	}

	return string(b)
}
