package leetcode

func FindMaxAverage(nums []int, k int) float64 {
	var maxSum int
	var curSum int
	for i := 0; i < k; i++ {
		curSum += int(nums[i])
	}
	maxSum = curSum
	for i := k; i < len(nums); i++ {
		curSum = curSum + int(nums[i]) - int(nums[i-k])
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return float64(maxSum) / float64(k)
}
