package dynamic

// func rob(nums []int) int {
// 	return help_rob(nums, len(nums)-1)
// }

// func help_rob(nums []int, i int) int {
// 	if i < 0 {
// 		return 0
// 	}
// 	return int(math.Max(float64(nums[i]+help_rob(nums, i-2)), float64(help_rob(nums, i-1))))
// }

func Rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prev1 := 0
	prev2 := 0
	for _, num := range nums {
		tmp := prev1
		prev1 = max(prev2+num, prev1)
		prev2 = tmp
	}
	return prev1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
