package leetcode

func oddCells(m int, n int, indices [][]int) int {
	matrix := makeMatrix(m, n)

	for _, index := range indices {
		addToMatrix(matrix, index[0], index[1])
	}

	return countOdd(matrix)

}

func makeMatrix(m, n int) [][]int {
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}
	return matrix
}

func addToMatrix(matrix [][]int, x, y int) {
	for i := 0; i < len(matrix); i++ {
		matrix[i][y]++
	}
	for j := 0; j < len(matrix[0]); j++ {
		matrix[x][j]++
	}
}

func countOdd(matrix [][]int) int {
	count := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j]%2 != 0 {
				count++
			}
		}
	}
	return count
}
