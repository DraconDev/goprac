package leetcode

func ScoreOfString(s string) int {
	score := 0
	for i, v := range s[0 : len(s)-1] {
		first := int(v - '0')
		second := int(s[i+1] - '0')
		score += abs(first - second)
	}
	return score
}
