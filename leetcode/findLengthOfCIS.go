package leetcode

func findLengthOfLCIS(nums []int) int {
	var longest int
	var sum int
	prev := 0

	for _, num := range nums {
		if num > prev {
			sum++
		} else {
			if sum > longest {
				longest = sum
			}
			sum = 1
		}

		prev = num
		if sum > longest {
			longest = sum
		}
	}
	return longest
}
