package tictactoe

func PlayTurn(x, y, turn int, board *[3][3]int) {
	board[x][y] = turn
}

func CheckWin(board [3][3]int) int {
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] && board[i][0] != 0 {
			return board[i][0]
		}
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] && board[0][i] != 0 {
			return board[0][i]
		}
	}
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[0][0] != 0 {
		return board[0][0]
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[0][2] != 0 {
		return board[0][2]
	}
	return 0
}

func CheckDraw(board [3][3]int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func Board(board [3][3]int) string {
	var out string
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			for j := 0; j < 5; j++ {
				if j%2 == 0 {
					if board[i/2][j/2] == 0 {
						out += " "
					}
					if board[i/2][j/2] == 1 {
						out += "X"
					}
					if board[i/2][j/2] == 2 {
						out += "O"
					}
				} else {
					out += "|"
				}
			}
		} else {
			out += "-+-+-"
		}
		out += "\n"
	}
	return out
}
