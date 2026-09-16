package main

func lengthOfLongestSubstring(s string) int {
	seen := [128]bool{}
	maxLength, left := 0, 0

	for right := 0; right < len(s); right++ {
		char := s[right]

		for seen[char] {
			seen[s[left]] = false
			left++
		}

		seen[char] = true

		if right-left+1 > maxLength {
			maxLength = right - left + 1
		}
	}

	return maxLength
}
