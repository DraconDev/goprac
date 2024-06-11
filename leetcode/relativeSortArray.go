package leetcode

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {

	existingElements := make(map[int]bool)
	for _, v := range arr2 {
		existingElements[v] = true
	}

	remainder := make([]int, 0)
	elements := make(map[int]int)
	for _, v := range arr1 {
		if existingElements[v] {
			elements[v]++
		} else {
			remainder = append(remainder, v)
		}
	}
	sort.Slice(remainder, func(i, j int) bool {
		return remainder[i] < remainder[j]
	})
	result := make([]int, 0)
	for _, v := range arr2 {
		for i := 0; i < elements[v]; i++ {
			result = append(result, v)
		}
	}
	result = append(result, remainder...)
	return result

}
