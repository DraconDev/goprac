package leetcode

func appendCharacters(s string, t string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		if result == len(t) {
			break
		}
		if s[i] == t[result] {
			result++
		}
	}
	return len(t) - result
}
