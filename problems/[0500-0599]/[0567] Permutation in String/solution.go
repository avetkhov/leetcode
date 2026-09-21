package main

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var seenS1, seenS2 [26]int

	for i := 0; i < len(s1); i++ {
		seenS1[s1[i]-'a']++
		seenS2[s2[i]-'a']++
	}

	if seenS1 == seenS2 {
		return true
	}

	for right := len(s1); right < len(s2); right++ {
		left := right - len(s1)

		seenS2[s2[right]-'a']++
		seenS2[s2[left]-'a']--

		if seenS1 == seenS2 {
			return true
		}
	}

	return false
}
