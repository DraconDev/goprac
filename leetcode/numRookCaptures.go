package leetcode

func numRookCaptures(board [][]byte) int {
	j, i := findRook(board)

	sum := 0
	for k := i; k >= 0; k-- {
		if board[j][k] == 'B' {
			break
		}
		if board[j][k] == 'p' {
			sum++
			break
		}

	}
	for k := i; k < len(board[i]); k++ {
		if board[j][k] == 'B' {
			break
		}
		if board[j][k] == 'p' {
			sum++
			break
		}
	}
	for k := j; k >= 0; k-- {
		if board[k][i] == 'B' {
			break
		}
		if board[k][i] == 'p' {
			sum++
			break
		}
	}
	for k := j; k < len(board[i]); k++ {
		if board[k][i] == 'B' {
			break
		}
		if board[k][i] == 'p' {
			sum++
			break
		}
	}

	return sum
}

func findRook(board [][]byte) (int, int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'R' {
				return i, j
			}
		}
	}
	return 0, 0
}
