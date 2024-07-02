package leetcode

func SubarraysWithKDistinct(nums []int, k int) int {
	return atMostKDistinct(nums, k) - atMostKDistinct(nums, k-1)
}

func atMostKDistinct(nums []int, k int) int {
	different := make(map[int]int)

	left := 0
	result := 0
	for i, v := range nums {
		different[v]++
		for len(different) > k {
			different[nums[left]]--
			if different[nums[left]] == 0 {
				delete(different, nums[left])
			}
			left++
		}
		result += i - left + 1
	}
	return result
}
