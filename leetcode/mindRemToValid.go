package leetcode

func minRemoveToMakeValid(s string) string {

	var brackets int

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			brackets++
		} else if s[i] == ')' {
			if brackets == 0 {
				s = s[:i] + s[i+1:]
				i--
			} else {
				brackets--
			}
		}
	}

	if brackets > 0 {
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == '(' {
				s = s[:i] + s[i+1:]
				brackets--
			}
			if brackets == 0 {
				break
			}
		}
	}
	return s

}
