package leetcode

import "sort"

func heightChecker(heights []int) int {
	ordered := make([]int, len(heights))
	copy(ordered, heights)
	sort.Ints(ordered)
	count := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != ordered[i] {
			count++
		}
	}
	return count
}
