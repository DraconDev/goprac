package leetcode

import "math"

func luckyNumbers(matrix [][]int) []int {
	RowsMin := make(map[int]int)
	for i := 0; i < len(matrix); i++ {
		min := math.MaxInt
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] < min {
				min = matrix[i][j]
			}
		}
		RowsMin[i] = min
	}

	ColumnsMax := make(map[int]int)
	for i := 0; i < len(matrix[0]); i++ {
		max := 0
		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] > max {
				max = matrix[j][i]
			}
		}
		ColumnsMax[max]++
	}

	result := []int{}

	for _, v := range RowsMin {
		if _, ok := ColumnsMax[v]; ok {
			result = append(result, v)
		}
	}
	return result
}
