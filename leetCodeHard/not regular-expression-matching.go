package leetCodeHard

import "strings"

// The matching should cover the entire input string (not partial).
/**
 	'.' Matches any single character.
	'*' Matches zero or more of the preceding element.
*/
/**
 * 匹配a、匹配a*、匹配.、匹配.*
 */

func regularExpressionMatching(s, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	if strings.Index(p, "**") >= 0 {
		return false
	}
	return false
}

// return index, length
func firstSimpleMatchIndex(s, p string) (int, int) { //a | . | a*
	if len(p) == 0 {
		return 0, 0
	}
	return 0, 0
}
