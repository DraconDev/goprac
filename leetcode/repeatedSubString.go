package leetcode

func RepeatedSubstringPattern(s string) bool {

	if len(s) < 2 {
		return false
	}

	sub := ""

	for i := 0; i < len(s)/2; i++ {
		sub += string(s[i])
		length := len(sub)
		for j := 0; j <= len(s)-length; j += length {
			if sub != s[j:j+length] {
				break
			}
			if j+length == len(s) {
				return true
			}
		}
	}

	return false

}
