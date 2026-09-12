package main

func isIsomorphic(s string, t string) bool {
	s2t, t2s := [256]byte{}, [256]byte{}

	for i := 0; i < len(s); i++ {
		chS, chT := s[i], t[i]

		if s2t[chS] != 0 && s2t[chS] != chT {
			return false
		}
		if t2s[chT] != 0 && t2s[chT] != chS {
			return false
		}

		s2t[chS], t2s[chT] = chT, chS
	}

	return true
}

func isIsomorphic(s string, t string) bool {
	s2t, t2s := make(map[byte]byte), make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		chS, chT := s[i], t[i]

		if mappedT, exists := s2t[chS]; exists && mappedT != chT {
			return false
		}
		if mappedS, exists := t2s[chT]; exists && mappedS != chS {
			return false
		}

		s2t[chS], t2s[chT] = chT, chS
	}

	return true
}
