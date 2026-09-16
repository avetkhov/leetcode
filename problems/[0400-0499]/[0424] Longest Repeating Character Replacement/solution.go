package main

func characterReplacement(s string, k int) int {
	seen := [26]int{}

	maxLength, maxFreq := 0, 0

	for left, right := 0, 0; right < len(s); right++ {
		seen[s[right]-'A']++

		if seen[s[right]-'A'] > maxFreq {
			maxFreq = seen[s[right]-'A']
		}

		if right-left+1-maxFreq > k {
			seen[s[left]-'A']--
			left++
		}

		if maxLength < right-left+1 {
			maxLength = right - left + 1
		}
	}

	return maxLength
}
