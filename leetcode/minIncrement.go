package leetcode

import "sort"

// func MinIncrementForUnique(nums []int) int {
// 	elements := make(map[int]int)
// 	for _, v := range nums {
// 		elements[v]++
// 	}
// 	increment := 0
// 	for k, v := range elements {
// 		if v == 1 {
// 			continue
// 		}
// 		counter := 0
// 		i := k
// 		for elements[k] > 1 {

// 			for elements[i] > 0 {
// 				counter++
// 				i++
// 			}
// 			elements[i]++
// 			increment += counter
// 			elements[k]--
// 		}
// 	}
// 	return increment
// }

func MinIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	increments := 0

	if len(nums) < 2 {
		return increments
	}

	curHigh := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] <= curHigh {
			increments += curHigh - nums[i] + 1
			curHigh = curHigh + 1
		} else {
			curHigh = nums[i]
		}
	}

	return increments

}
