package leetcode

func backspaceCompare(s string, t string) bool {
	return string(backspaceConvert(s)) == string(backspaceConvert(t))
}

func backspaceConvert(s string) []byte {
	result := []byte{}
	skip := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '#' {
			skip++
			continue
		}
		if skip > 0 {
			skip--
		} else {
			result = append(result, s[i])
		}
	}
	return result
}
