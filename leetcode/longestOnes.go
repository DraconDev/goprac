package leetcode

func LongestOnes(nums []int, k int) int {
	start := 0
	end := 0
	for end < len(nums) {
		if nums[end] == 0 {
			k--
		}
		if k < 0 {
			if nums[start] == 0 {
				k++
			}
			start++
		}
		end++
	}
	return end - start
}
