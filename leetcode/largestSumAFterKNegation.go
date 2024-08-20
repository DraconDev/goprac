package leetcode

import "sort"

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	position := 0
	for k != 0 {
		if nums[position] == 0 {
			break
		}
		nums[position] = -nums[position]
		k--
		if position+1 < len(nums) && abs(nums[position]) > nums[position+1] {
			position++
		}
	}
	return sum(nums)
}

func sum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// func abs(a int) int {
// 	if a < 0 {
// 		return -a
// 	}
// 	return a
// }
