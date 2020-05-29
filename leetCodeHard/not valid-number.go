package leetCodeHard

//TODO: test case 不完整
func isNumber(s string) bool {
	return false
	// preR := []byte{"-"[0], "+"[0]}
	// isFloat := false
	// afterE := false
	// hasNum := false
	// s = strings.TrimSpace(s)
	// if len(s) <= 0 {
	// 	return false
	// }
	// for i, _ := range s {
	// 	v := s[i]
	// 	shouContinue := false
	// 	for _, vv := range preR {
	// 		if vv == v {
	// 			if i == 0 {
	// 				shouContinue = true
	// 				break
	// 			} else if vv == "-"[0] && s[i-1] == "e"[0] {
	// 				shouContinue = true
	// 				break
	// 			} else {
	// 				return false
	// 			}
	// 		}
	// 	}
	// 	if shouContinue {
	// 		continue
	// 	}

	// 	if v >= "0"[0] && v <= "9"[0] {
	// 		hasNum = true
	// 		continue
	// 	}

	// 	if v == "."[0]  {
	// 		if isFloat {
	// 			return false
	// 		}
	// 		if i == 0 || !hasNum {
	// 			if len(s) > i+1 {
	// 				isFloat = true
	// 				continue
	// 			}
	// 			return false
	// 		}

	// 		if !isFloat && !afterE {
	// 			isFloat = true
	// 			continue
	// 		}
	// 		return false
	// 	}

	// 	if v == "e"[0] && len(s) > i+1 {
	// 		if !afterE && hasNum {
	// 			afterE = true
	// 			continue
	// 		}
	// 		return false
	// 	}
	// 	return false
	// }
	// return true
}
