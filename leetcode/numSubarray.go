package leetcode

func NumSubarraysWithSum(nums []int, goal int) int {

	var sum int
	var curSum int

	for i, _ := range nums {
		for _, v2 := range nums[i:] {
			curSum += v2
			if curSum == goal {
				sum++
				curSum = 0
			}

		}
	}

	return sum

}
