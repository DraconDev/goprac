package leetcode

func findMaxConsecutiveOnes(nums []int) int {
	var max int
	currMax := 0
	for _, v := range nums {
		if v == 1 {
			currMax++
		} else {
			currMax = 0
		}
		if currMax > max {
			max = currMax
		}
	}
	return max
}
