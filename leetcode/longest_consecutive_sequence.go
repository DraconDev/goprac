package leetcode

import "sort"

func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	result := 1
	temp := 1
	for i, v := range nums[1:] {
		if v == nums[i] {
			continue
		}
		if v == nums[i]+1 {
			temp++
			if temp > result {
				result = temp
			}
		} else {
			temp = 1
		}
	}
	return result
}
