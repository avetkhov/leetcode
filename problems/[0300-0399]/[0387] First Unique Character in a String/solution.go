package main

// UTF-8
func firstUniqChar(s string) int {
	count := make(map[rune]int)

	for _, ch := range s {
		count[ch]++
	}

	for i, ch := range s {
		if count[ch] == 1 {
			return i
		}
	}

	return -1
}

// ASCII
func firstUniqChar(s string) int {
	count := [26]int{}

	for _, ch := range s {
		count[ch-'a']++
	}

	for i, ch := range s {
		if count[ch-'a'] == 1 {
			return i
		}
	}

	return -1
}
