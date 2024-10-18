package leetcode

func tictactoe(moves [][]int) string {

	if len(moves) < 5 {
		return "Pending"
	}

	a := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	b := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	for i := 0; i < len(moves); i++ {
		if i%2 == 0 {
			a[moves[i][0]][moves[i][1]] = 1
		} else {
			b[moves[i][0]][moves[i][1]] = 1
		}
	}

	if checkWin(a) {
		return "A"
	} else if checkWin(b) {
		return "B"
	} else if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}

func checkWin(player [][]int) bool {
	for i := 0; i < 3; i++ {
		if player[i][0] == player[i][1] && player[i][0] == player[i][2] && player[i][0] != 0 {
			return true
		}
	}
	for i := 0; i < 3; i++ {
		if player[0][i] == player[1][i] && player[0][i] == player[2][i] && player[0][i] != 0 {
			return true
		}
	}
	if player[0][0] == player[1][1] && player[0][0] == player[2][2] && player[0][0] != 0 {
		return true
	}
	if player[0][2] == player[1][1] && player[0][2] == player[2][0] && player[0][2] != 0 {
		return true
	}
	return false
}
