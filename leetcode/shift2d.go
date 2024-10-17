package leetcode

func shiftGrid(grid [][]int, k int) [][]int {

	m, n := len(grid), len(grid[0])
	copy := make([][]int, m)
	for i := 0; i < m; i++ {
		copy[i] = make([]int, n)
	}

	k = k % (m * n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			idx := (i*n + j + k) % (m * n)
			copy[idx/n][idx%n] = grid[i][j]
		}
	}
	return copy

}
