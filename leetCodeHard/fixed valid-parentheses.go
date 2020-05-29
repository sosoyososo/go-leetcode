package leetCodeHard

/**
 * {} [] () 判断
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.
*/
func isValidParentheses(s string) bool {
	matchMap := map[byte]byte{
		[]byte("(")[0]: []byte(")")[0],
		[]byte("[")[0]: []byte("]")[0],
		[]byte("{")[0]: []byte("}")[0],
	}
	ret := []byte{}
	for i, _ := range s {
		if s[i] == []byte("(")[0] || s[i] == []byte("[")[0] || s[i] == []byte("{")[0] {
			ret = append(ret, s[i])
			continue
		}

		if s[i] == []byte(")")[0] || s[i] == []byte("]")[0] || s[i] == []byte("}")[0] {
			if len(ret) <= 0 || matchMap[ret[len(ret)-1]] != s[i] {
				return false
			}
			ret = ret[:len(ret)-1]
		}
	}

	if len(ret) > 0 {
		return false
	}
	return true
}
