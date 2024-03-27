package leetcode

func numSubarrayProductLessThanK(nums []int, k int) int {
	result := 0

	for i := 0; i < len(nums); i++ {
		temp := 1
		for j := i; j < len(nums); j++ {
			temp *= nums[j]
			if temp < k {
				result++
			} else {
				break
			}
		}
	}

	return result

}
