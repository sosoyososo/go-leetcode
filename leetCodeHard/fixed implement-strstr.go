package leetCodeHard

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}

	for i, v := range haystack {
		if v != rune(needle[0]) {
			continue
		}

		match := true
		for j := 0; j < len(needle); j++ {
			if i+len(needle) > len(haystack) {
				match = false
				break
			}
			if haystack[i+j] != needle[j] {
				match = false
				break
			}
		}

		if match {
			return i
		}
	}
	return -1
}
