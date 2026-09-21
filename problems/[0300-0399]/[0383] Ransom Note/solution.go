package main

func canConstruct(ransomNote string, magazine string) bool {
	var seen [26]int

	for i := 0; i < len(magazine); i++ {
		seen[magazine[i]-'a']++
	}

	for i := 0; i < len(ransomNote); i++ {
		idx := ransomNote[i] - 'a'
		seen[idx]--

		if seen[idx] < 0 {
			return false
		}
	}

	return true
}
