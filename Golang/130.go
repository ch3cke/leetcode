package main

func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])
	result := make([][]byte, m)
	for i := 0; i < m; i++ {
		result[i] = make([]byte, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = 'X'
		}
	}
	i := 0
	j := 0
	var nextNode func(i, j int)
	nextNode = func(i, j int) {
		if i < 0 || i >= m {
			return
		}
		if j < 0 || j >= n {
			return
		}
		if board[i][j] == 'O' && result[i][j] == 'X' {
			result[i][j] = 'O'
			nextNode(i+1, j)
			nextNode(i-1, j)
			nextNode(i, j-1)
			nextNode(i, j+1)
		} else {
			return
		}
	}

	for ; j < n; j++ {
		if board[0][j] == 'O' {
			nextNode(0, j)
		}
		if board[m-1][j] == 'O' {
			nextNode(m-1, j)
		}
	}
	for ; i < m; i++ {
		if board[i][0] == 'O' {
			nextNode(i, 0)
		}
		if board[i][n-1] == 'O' {
			nextNode(i, n-1)
		}
	}
	board = result
}
