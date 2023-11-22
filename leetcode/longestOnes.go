package leetcode

func LongestOnes(nums []int, k int) int {
	start := 0
	end := 0
	for end < len(nums) {
		k -= 1 - nums[end]
		if k < 0 {
			k += 1 - nums[start]
			start++
		}
		end++
	}
	return end - start
}
