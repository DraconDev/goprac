package leetcode

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	var result int
	for i := 0; i < len(nums); i += 2 {

		result += nums[i]
	}
	return result

}
