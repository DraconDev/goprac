package leetcode

// func LongestSubarrayAfterDel(nums []int) int {
// 	max := 0
// 	var left, right, leftLag = 0, 0, 0
// 	for right < len(nums) {
// 		if nums[right] == 0 {
// 			currentMax := right - left - 1
// 			if currentMax > max {
// 				max = currentMax
// 			}
// 			left = leftLag
// 			leftLag = right + 1
// 		}
// 		right++
// 	}
// 	if right-left-1 > max {
// 		max = right - left - 1
// 	}
// 	return max
// }

func LongestSubarrayAfterDel(nums []int) int {
	n := len(nums)

	left := 0
	zeros := 0
	ans := 0

	for right := 0; right < n; right++ {
		if nums[right] == 0 {
			zeros++
		}

		for zeros > 1 {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}

		ans = max(ans, right-left+1-zeros)
	}

	if ans == n {
		return ans - 1
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}