package leetcode

func FindMaxAverage(nums []int, k int) float64 {
	var maxSum float64
	var curSum float64
	for i := 0; i < k; i++ {
		curSum += float64(nums[i])
	}
	maxSum = curSum
	for i := k; i < len(nums); i++ {
		curSum = curSum + float64(nums[i]) - float64(nums[i-k])
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum / float64(k)
}
