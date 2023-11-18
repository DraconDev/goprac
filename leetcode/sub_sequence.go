package leetcode

// func IsSubsequence(s string, t string) bool {
// 	if len(s) == 0 {
// 		return true
// 	}
// 	for i, v := range t {
// 		if v == rune(s[0]) {
// 			for j := i + 1; j < len(t); j++ {
// 				if t[j] == s[1] {
// 					for k := j + 1; k < len(t); k++ {
// 						if t[k] == s[2] {
// 							return true
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return false
// }

func IsSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	for i, v := range t {
		if v == rune(s[0]) {
			return IsSubsequence(s[1:], t[i+1:])
		}
	}
	return false
}

// func findElem(s string, t string, loop int) bool {
// 	for i, v := range t {
// 		if v == rune(s[0]) {
// 			for j := i + 1; j < len(t); j++ {
// 				if t[j] == s[1] {
// 					return true
// 				}
// 			}
// 		}
// 	}
// 	return false
// }
