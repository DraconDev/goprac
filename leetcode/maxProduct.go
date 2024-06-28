package leetcode

import "sort"

func maximumProduct(nums []int) int {
	// sort descending
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	return max(nums[0]*nums[1]*nums[2], nums[0]*nums[len(nums)-1]*nums[len(nums)-2])
}
