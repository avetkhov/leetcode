package main

import "strconv"

func compress(chars []byte) int {
	r, k, n := 0, 0, len(chars)
	for r < n {
		l := r + 1
		for l < n && chars[r] == chars[l] {
			l++
		}

		chars[k] = chars[r]
		k++

		count := l - r
		if count > 1 {
			countStr := strconv.Itoa(count)
			for _, v := range countStr {
				chars[k] = byte(v)
				k++
			}
		}

		r = l
	}

	return k
}
