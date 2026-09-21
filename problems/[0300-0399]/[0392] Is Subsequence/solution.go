package main

func isSubsequence(s string, t string) bool {
	pS, pT := 0, 0

	for pS < len(s) && pT < len(t) {
		if s[pS] == t[pT] {
			pS++
		}

		pT++
	}

	return pS == len(s)
}
