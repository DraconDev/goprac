package leetcode

import "sort"

func frequencySort(nums []int) []int {

	if len(nums) == 0 {
		return nums
	}

	frequencyMap := make(map[int]int)
	for _, v := range nums {
		frequencyMap[v]++
	}

	maxFreq := len(nums)
	items := make([][]int, maxFreq+1)
	for num, freq := range frequencyMap {
		items[freq] = append(items[freq], num)
	}

	result := []int{}

	for i := 0; i < len(items); i++ {
		if len(items[i]) == 0 {
			continue
		}
		sort.Sort(sort.Reverse(sort.IntSlice(items[i])))
		for _, v := range items[i] {
			for j := 0; j < i; j++ {
				result = append(result, v)
			}
		}
	}
	return result

}
