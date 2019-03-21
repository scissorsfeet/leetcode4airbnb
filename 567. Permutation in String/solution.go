package _67__Permutation_in_String

func checkInclusion(s1 string, s2 string) bool {
	s1Len := len(s1)
	s1Map := [26]rune{}
	for _, b := range s1 {
		s1Map[b-'a']++
	}

	s2Map := [26]rune{}
	for i, b := range s2 {
		if i+1-s1Len >= 0 {
			if i+1-s1Len > 0 {
				s2Map[s2[i-s1Len]-'a']--
			}
		}
		s2Map[b-'a']++

		if isSame(s1Map, s2Map) {
			return true
		}
	}

	return false
}

func isSame(a, b [26]rune) bool {
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
