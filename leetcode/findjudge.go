package leetcode

func FindJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}

	votes := map[int]int{}
	trusting := map[int]int{}
	for _, e := range trust {
		votes[e[1]]++
		trusting[e[0]]++
	}

	for k, v := range votes {
		if v == n-1 && trusting[k] == 0 {
			return k
		}
	}
	return -1

}
