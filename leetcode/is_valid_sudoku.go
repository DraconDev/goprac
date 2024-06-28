package leetcode

func IsValidSudoku(board [][]byte) bool {
	// * check rows
	for i := 0; i < 9; i++ {
		row := make(map[byte]int)
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			row[board[i][j]]++
			if row[board[i][j]] > 1 {
				return false
			}
		}
	}
	// * check columns
	for i := 0; i < 9; i++ {
		column := make(map[byte]int)
		for j := 0; j < 9; j++ {
			if board[j][i] == '.' {
				continue
			}
			column[board[j][i]]++
			if column[board[j][i]] > 1 {
				return false
			}
		}
	}
	// * check boxes
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			box := make(map[byte]int)
			for x := i; x < i+3; x++ {
				for y := j; y < j+3; y++ {
					if board[x][y] == '.' {
						continue
					}
					box[board[x][y]]++
					if box[board[x][y]] > 1 {
						return false
					}
				}
			}
		}
	}

	return true
}
