package main

func reverseWords(s string) string {
	b := []byte(s)

	for l, r := 0, 0; r <= len(b); r++ {
		if r == len(b) || b[r] == ' ' {
			for i, j := l, r-1; i < j; i, j = i+1, j-1 {
				b[i], b[j] = b[j], b[i]
			}

			l = r + 1
		}
	}

	return string(b)
}
