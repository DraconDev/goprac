package leetcode

func MinimumLength(s string) int {

	for len(s) > 1 && s[0] == s[len(s)-1] {

		start := 1
		end := len(s) - 1
		for i := 1; i < end; i++ {
			if s[i] == s[0] {
				start++
			} else {
				break
			}
		}

		for i := len(s) - 2; i > start; i-- {
			if s[i] == s[0] {
				end--
			} else {
				break
			}
		}

		s = s[start:end]
	}
	return len(s)

}
