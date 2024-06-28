package leetcode

import "math"

func MinFallingPathSum(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	for r := 1; r < m; r++ {
		for c := 0; c < n; c++ {
			switch c {
			case 0:
				matrix[r][c] += Min(matrix[r-1][c], matrix[r-1][c+1])
			case n - 1:
				matrix[r][c] += Min(matrix[r-1][c-1], matrix[r-1][c])
			default:
				matrix[r][c] += Min(matrix[r-1][c-1], Min(matrix[r-1][c], matrix[r-1][c+1]))
			}
		}
	}
	min := math.MaxInt
	for c := 0; c < n; c++ {
		min = Min(min, matrix[m-1][c])
	}
	return min
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
