package leetcode

func numEquivDominoPairs(dominoes [][]int) int {

	matches := map[[2]int]int{}

	eq := 0

	for _, domino := range dominoes {

		if domino[0] > domino[1] {
			domino[0], domino[1] = domino[1], domino[0]
		}
		d := [2]int{domino[0], domino[1]}
		matches[d]++

		if matches[d] > 1 {
			eq += matches[d] - 1
		}
	}
	return eq
}
