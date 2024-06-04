package leetcode

func CountBinarySubstrings(s string) int {
	prev, cur, res := 0, 1, 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cur++
		} else {
			prev = cur
			cur = 1
		}
		if prev >= cur {
			res++
		}
	}
	return res
}
