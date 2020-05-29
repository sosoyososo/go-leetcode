package leetCodeHard

import "fmt"

/*
	最长回文字符串
*/
func longestPalindrome(s string) string {
	if len(s) <= 0 {
		return ""
	}

	/*
		找到回文点， return index, isOdd
	*/
	findPointFromIndex := func(index int) int {
		for i := index; i < len(s); i++ {
			if i+1 < len(s) && s[i] == s[i+1] {
				return i
			} else if i-1 >= 0 && i+1 < len(s) && s[i-1] == s[i+1] {
				return i
			}
		}
		return -1
	}

	palindromeStrAtPoint := func(index int, isOdd bool) string {
		ret := ""
		pre := ""
		l := len(s)
		for i := 0; ; i++ {
			if isOdd {
				if i == 0 {
				} else if index-i >= 0 && index+i < l {
					if s[index-i] != s[index+i] {
						break
					}
				} else {
					break
				}
			} else {
				if index-i >= 0 && index+i+1 < l {
					if s[index-i] != s[index+i+1] {
						break
					}
				} else {
					break
				}
			}
			ret += string(s[index-i])
			if isOdd && i == 0 {
				continue
			}
			pre = string(s[index-i]) + pre
		}
		return pre + ret
	}

	maxRet := ""
	nextFindStartIndex := 0
	for {
		index := findPointFromIndex(nextFindStartIndex)
		fmt.Printf("%v\n", index)
		if index >= 0 {
			str := palindromeStrAtPoint(index, false)
			str2 := palindromeStrAtPoint(index, true)
			if len(str2) > len(str) {
				str = str2
			}
			if len(str) > len(maxRet) {
				maxRet = str
			}
			nextFindStartIndex = index + 1
		} else {
			break
		}
	}

	if len(maxRet) <= 0 {
		maxRet = string(s[0])
	}

	return maxRet
}
