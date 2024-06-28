package leetcode

import "sort"

func LargestPerimeter(nums []int) int64 {

	if len(nums) < 3 {
		return -1
	}

	sort.Ints(nums)

	var result int

	for _, v := range nums {
		result += int(v)
	}

	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i] >= result-nums[i] {
			result -= nums[i]
		} else {
			return int64(result)
		}
	}

	return -1

}
