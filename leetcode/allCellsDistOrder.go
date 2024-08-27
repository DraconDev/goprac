package leetcode

import "sort"

func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {

	var result [][]int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result = append(result, []int{i, j})
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return abs(result[i][0]-rCenter)+abs(result[i][1]-cCenter) < abs(result[j][0]-rCenter)+abs(result[j][1]-cCenter)
	})
	return result
}

