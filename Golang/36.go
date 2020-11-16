package main

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		res := [3][9]bool{}
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				col := board[i][j] % '1'
				if res[0][col] {
					return false
				}
				res[0][col] = true
			}
			if board[j][i] != '.' {
				row := board[i][j] % '1'
				if res[1][row] {
					return false
				}
				res[1][row] = true
			}
			if board[j/3+i%3*3][i/3*3+j%3] != '.' {
				saw := board[j/3+i%3*3][i/3*3+j%3] % '1'
				if res[2][saw] {
					return false
				}
				res[2][saw] = true
			}
		}
	}
	return true
}
