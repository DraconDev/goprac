package leetcode

func surfaceArea(grid [][]int) int {
	var area int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] > 0 {
				area += 2 // top and bottom
				if i > 0 {
					area += max(0, grid[i][j]-grid[i-1][j])
				} else {
					area += grid[i][j]
				}
				if i < len(grid)-1 {
					area += max(0, grid[i][j]-grid[i+1][j])
				} else {
					area += grid[i][j]
				}
				if j > 0 {
					area += max(0, grid[i][j]-grid[i][j-1])
				} else {
					area += grid[i][j]
				}
				if j < len(grid[0])-1 {
					area += max(0, grid[i][j]-grid[i][j+1])
				} else {
					area += grid[i][j]
				}
			}
		}
	}
	return area
}
