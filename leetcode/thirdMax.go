package leetcode

import "sort"

func ThirdMax(nums []int) int {

	// filter out dups
	seen := make(map[int]bool)
	result := make([]int, 0)
	for _, v := range nums {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	// sort descending
	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})

	length := len(result)
	if length >= 3 {
		return result[2]
	}
	return result[0]
}
