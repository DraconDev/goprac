package leetcode

func LongestSubarrayAfterDel(nums []int) int {
	max := 0
	var left, right, leftLag = 0, 0, 0
	for right < len(nums) {
		if nums[right] == 0 {
			currentMax := right - left - 1
			if currentMax > max {
				max = currentMax
			}
			left = leftLag
			leftLag = right + 1
		}
		right++
	}
	if right-left-1 > max {
		max = right - left - 1
	}
	return max
}
