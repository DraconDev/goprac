package main

// func maxSubArray(nums []int) int {
// 	if len(nums) == 1 {
// 		return nums[0]
// 	}
// 	index := 1
// 	maxSum := nums[0]
// 	currentSub := nums[0]
// 	for index < len(nums) {
// 		currentSub += nums[index]
// 		if currentSub > maxSum && currentSub > nums[index] {
// 			maxSum = currentSub
// 		} else if nums[index] > maxSum {
// 			maxSum = nums[index]
// 			currentSub = nums[index]
// 		}
// 		if currentSub < 0 {
// 			currentSub = 0
// 		}
// 		index++
// 	}
// 	return maxSum
// }

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum := nums[0]
	currentSub := 0
	for _, e := range nums {
		currentSub += e
		if currentSub > maxSum {
			maxSum = currentSub
		}
		if currentSub < 0 {
			currentSub = 0
		}
	}
	return maxSum
}
