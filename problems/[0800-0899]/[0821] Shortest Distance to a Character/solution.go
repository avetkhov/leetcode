package main

func shortestToChar(s string, c byte) []int {
	result := make([]int, len(s))

	pos := -2 * len(s)
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			pos = i
		}

		result[i] = i - pos
	}

	pos = 2 * len(s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			pos = i
		}

		if pos-i < result[i] {
			result[i] = pos - i
		}
	}

	return result
}
