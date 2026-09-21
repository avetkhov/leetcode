package main

func maxVowels(s string, k int) int {
	maxCount := 0

	for i := 0; i < k; i++ {
		maxCount += isVowel(s[i])
	}

	count := maxCount
	for i := k; i < len(s); i++ {
		count += isVowel(s[i]) - isVowel(s[i-k])
		if count > maxCount {
			maxCount = count
		}

		if maxCount == k {
			return maxCount
		}
	}

	return maxCount
}

func isVowel(b byte) int {
	switch b {
	case 'a', 'e', 'i', 'o', 'u':
		return 1
	default:
		return 0
	}
}
