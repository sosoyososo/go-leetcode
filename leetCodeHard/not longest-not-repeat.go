package leetCodeHard

func longestNotRepeat(s string) string {
	if len(s) <= 0 {
		return ""
	}

	maxRet := []byte{}

	ret := []byte{}
	findIndex := func(r byte) int {
		for i, v := range ret {
			if r == v {
				return i
			}
		}
		return -1
	}

	for i := 0; i < len(s); i++ {
		r := s[i]
		if findIndex(r) < 0 {
			ret = append(ret, r)
		} else {
			if len(maxRet) < len(ret) {
				maxRet = ret
				i -= len(ret)
			}
		}
	}

	return string(maxRet)
}
