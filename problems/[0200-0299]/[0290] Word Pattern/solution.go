package main

import "strings"

func wordPattern(pattern string, s string) bool {
	strs := strings.Fields(s)

	if len(pattern) != len(strs) {
		return false
	}

	charToWord := make(map[byte]string)
	wordToChar := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		ch, str := pattern[i], strs[i]

		if mappedWord, exists := charToWord[ch]; exists && mappedWord != str {
			return false
		}

		if mappedChar, exists := wordToChar[str]; exists && mappedChar != ch {
			return false
		}

		charToWord[ch], wordToChar[str] = str, ch
	}

	return true
}
