package leetcode

import "sort"

func findMaxK(nums []int) int {
	sort.Ints(nums)

	left := 0
	right := len(nums) - 1

	for left < right {
		if nums[left]+nums[right] < 0 {
			left++
		} else if nums[left]+nums[right] > 0 {
			right--
		} else {
			return nums[right]
		}
	}
	return -1
}
