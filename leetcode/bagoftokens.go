package leetcode

import "sort"

func BagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)

	var score int

	for len(tokens) > 0 && power >= tokens[0] {
		power -= tokens[0]
		tokens = tokens[1:]
		score += 1

		if len(tokens) > 1 && tokens[0] > power {
			power += tokens[len(tokens)-1]
			tokens = tokens[:len(tokens)-1]
			score -= 1
		}
	}

	return score
}
