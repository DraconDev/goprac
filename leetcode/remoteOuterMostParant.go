package leetcode

func removeOuterParentheses(s string) string {

	var result string
	var count int
	for _, v := range s {
		if v == '(' {
			if count > 0 {
				result += string(v)
			}
			count++
		} else {
			count--
			if count > 0 {
				result += string(v)
			}
		}
	}
	return result
}
