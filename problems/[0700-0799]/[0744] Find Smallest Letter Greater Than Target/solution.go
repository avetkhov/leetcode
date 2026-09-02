package main

func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)-1
	a := letters[0]

	for l <= r {
		m := l + (r-l)/2

		if letters[m] > target {
			a = letters[m]
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return a
}
