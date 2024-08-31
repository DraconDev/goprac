package leetcode

import "sort"

func lastStoneWeight(stones []int) int {

	if len(stones) == 0 {
		return 0
	}

	if len(stones) == 1 {
		return stones[0]
	}

	for len(stones) > 1 {
		sort.Ints(stones)
		if stones[len(stones)-1] == stones[len(stones)-2] {
			stones = stones[:len(stones)-2]
			continue
		}
		stones[len(stones)-2] = stones[len(stones)-1] - stones[len(stones)-2]
		stones = stones[:len(stones)-1]
	}
	if len(stones) == 0 {
		return 0
	}

	return stones[0]
}
