package leetcode

func MaxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		maximum := 0
		maxSum := 0
		for j := 1; j <= k && i-j >= 0; j++ {
			maximum = max(maximum, arr[i-j])
			maxSum = max(maxSum, dp[i-j]+maximum*j)
		}
		dp[i] = maxSum
	}

	return dp[n]
}
