package leetcode

import "strconv"

func findNumbers(nums []int) int {
	result := 0

	for _, v := range nums {
		if len(strconv.Itoa(v))%2 == 0 {
			result++
		}
	}
	return result
}
