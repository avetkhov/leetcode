package main

func isPerfectSquare(num int) bool {
	l, r := 1, num

	for l <= r {
		m := l + (r-l)/2
		s := m * m

		if s == num {
			return true
		} else if s > num {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return false
}
