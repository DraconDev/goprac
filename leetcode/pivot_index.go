package leetcode

func PivotIndex(nums []int) int {
	var sum = 0
	for _, v := range nums {
		sum += v
	}
	var leftSum = 0
	for i, v := range nums {
		if leftSum == sum-leftSum-v {
			return i
		}
		leftSum += v
	}
	return -1
}
