package main

func backspaceCompare(s string, t string) bool {
	i, j := len(s)-1, len(t)-1

	count := 0
	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				count++
				i--
			} else if count > 0 {
				count--
				i--
			} else {
				break
			}

		}

		count = 0
		for j >= 0 {
			if t[j] == '#' {
				count++
				j--
			} else if count > 0 {
				count--
				j--
			} else {
				break
			}
		}

		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if (i >= 0) != (j >= 0) {
			return false
		}

		i--
		j--
	}

	return true
}
