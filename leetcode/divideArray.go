package leetcode

import "sort"

func DivideArray(nums []int, k int) [][]int {
	if len(nums)%3 != 0 {
		return [][]int{}
	}
	sort.Ints(nums)

	var result [][]int

	// one or more size 3
	var temp []int
	for i, e := range nums {

		temp = append(temp, e)

		if (i+1)%3 == 0 {
			if e > temp[0]+k {
				return [][]int{}
			}
			result = append(result, temp)
			temp = []int{}
		}
	}
	return result
}
