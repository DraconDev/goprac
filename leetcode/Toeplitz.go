package leetcode

func IsToeplitzMatrix(matrix [][]int) bool {
	for i := 0; i < len(matrix[0]); i++ {
		for j := 1; j < min(len(matrix), len(matrix[0])-i); j++ {
			if matrix[0][i] != matrix[j][i+j] {
				return false
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < min(len(matrix[0]), len(matrix)-i); j++ {
			if matrix[i][0] != matrix[i+j][j] {
				return false
			}
		}
	}

	return true
}
