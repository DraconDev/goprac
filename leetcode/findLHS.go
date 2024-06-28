package leetcode

import "sort"

func findLHS(nums []int) int {
	sort.Ints(nums)

	// remove duplicates
	uniqueNums := make([]int, 0)
	seen := make(map[int]bool)
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
		if !seen[v] {
			seen[v] = true
			uniqueNums = append(uniqueNums, v)
		}
	}

	max := 0
	for _, v := range uniqueNums[:len(uniqueNums)-1] {
		if count[v+1] > 0 {
			if count[v+1]+count[v] > max {
				max = count[v+1] + count[v]
			}
		}
	}

	return max

}
