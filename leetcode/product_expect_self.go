package leetcode

func ProductExceptSelf(nums []int) []int {
	var pre, post = 1, 1
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = pre
		pre *= nums[i]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= post
		post *= nums[i]
	}
	return result
}
