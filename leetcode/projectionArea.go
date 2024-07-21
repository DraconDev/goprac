package leetcode

// func projectionArea(grid [][]int) int {
// 	result := 0

// 	for i := 0; i < len(grid); i++ {
// 		for j := 0; j < len(grid[0]); j++ {
// 			if grid[i][j] > 0 {
// 				result++
// 			}
// 		}
// 	}

// 	for i, _ := range grid[0] {
// 		highest := 0
// 		for j := 0; j < len(grid[0]); j++ {
// 			if grid[j][i] > highest {
// 				highest = grid[j][i]
// 			}
// 		}
// 		result += highest
// 	}

// 	for i := 0; i < len(grid); i++ {
// 		highest := 0
// 		for j := 0; j < len(grid[0]); j++ {
// 			if grid[i][j] > highest {
// 				highest = grid[i][j]
// 			}
// 		}
// 		result += highest
// 	}
// 	return result
// }

func projectionArea(grid [][]int) int {
	result := 0
	for i, _ := range grid[0] {
		highestX := 0
		highestY := 0
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] > 0 {
				result++
			}
			if grid[i][j] > highestX {
				highestX = grid[i][j]
			}

			if grid[j][i] > highestY {
				highestY = grid[j][i]
			}
		}
		result += highestX
		result += highestY
	}
	return result
}
