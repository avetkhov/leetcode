package main

func isHappy(n int) bool {
	s, f := n, getNext(n)

	for f != 1 && s != f {
		s, f = getNext(s), getNext(getNext(f))
	}

	return f == 1
}

func getNext(n int) int {
	sum := 0
	for ; n > 0; n /= 10 {
		sum += (n % 10) * (n % 10)
	}
	return sum
}
