package leetcode

func LargestLocal(grid [][]int) [][]int {

	result := [][]int{}
	var max int

	for i := 0; i < len(grid)-2; i++ {
		row := []int{}
		for j := 0; j < len(grid)-2; j++ {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					if grid[i+k][j+l] > max {
						max = grid[i+k][j+l]
					}
				}
			}
			row = append(row, max)
			max = 0
		}
		result = append(result, row)
	}
	return result
}
