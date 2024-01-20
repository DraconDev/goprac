package leetcode

import "math"

func SumSubarrayMins(arr []int) int {
	var result int

	for i := 0; i < len(arr); i++ {
		min := math.MaxInt
		for j := i; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
			}
			result += min
		}
	}

	return result

}
