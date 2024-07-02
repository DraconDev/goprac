package leetcode

import (
	"sort"
)

func sortedSquares(nums []int) []int {
	for i, v := range nums {
		if v < 0 {
			v = -v
		}
		nums[i] = v * v
	}

	sort.Ints(nums)
	return nums

}
