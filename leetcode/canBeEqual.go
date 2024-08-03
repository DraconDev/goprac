package leetcode

import "sort"

func canBeEqual(target []int, arr []int) bool {

	if len(target) == 0 || len(target) != len(arr) {
		return false
	}

	sort.Ints(target)
	sort.Ints(arr)

	for i, v := range target {
		if v != arr[i] {
			return false
		}
	}
	return true
}
