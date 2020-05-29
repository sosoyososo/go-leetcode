package leetCodeHard

func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{}
	}

	strs := generate("", n)
	for {
		tmp := []string{}
		for _, str := range strs {
			list := generate(str, n)
			for _, v := range list {
				tmp = append(tmp, str+v)
			}
		}
		strs = tmp
		if len(strs[0]) == 2*n {
			break
		}
	}
	return strs
}

func generate(str string, t int) []string {
	c1 := 0
	c2 := 0
	for _, v := range str {
		if v == rune("("[0]) {
			c1++
		} else {
			c2++
		}
	}
	if c2 == c1 {
		return []string{"("}
	} else if c1 < t {
		return []string{")", "("}
	}
	return []string{")"}
}
