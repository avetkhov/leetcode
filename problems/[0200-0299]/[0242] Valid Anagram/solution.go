package main

// UTF-8
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[rune]int, len(s))

	for _, letter := range s {
		count[letter]++
	}

	for _, letter := range t {
		count[letter]--
		if count[letter] < 0 {
			return false
		}
	}

	return true
}

// ASCII
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := [26]int{}

	for i := range s {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}

	for _, c := range count {
		if c != 0 {
			return false
		}
	}

	return true
}
