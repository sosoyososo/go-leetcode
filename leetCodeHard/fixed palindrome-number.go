package leetCodeHard

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	str := strconv.Itoa(x)
	for i := 0; ; i++ {
		s := i
		e := len(str) - i - 1
		if s >= e {
			break
		}
		if str[s] != str[e] {
			return false
		}
	}
	return true
}
