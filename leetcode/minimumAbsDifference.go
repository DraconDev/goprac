package leetcode

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)

	min := math.MaxInt

	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] < min {
			min = arr[i+1] - arr[i]
		}
	}
	result := [][]int{}

	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] == min {
			result = append(result, []int{arr[i], arr[i+1]})
		}
	}
	return result
}
