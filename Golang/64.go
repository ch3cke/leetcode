package main

func minPathSum(grid [][]int) int {
	if len(grid[0]) == 0 || len(grid) == 0 {
		return 0
	}
	cloun, row := len(grid[0]), len(grid)
	tmp := make([][]int, row)
	for i := 0; i < row; i++ {
		tmp[i] = make([]int, cloun)
	}
	tmp[0][0] = grid[0][0]
	for i := 1; i < cloun; i++ {
		tmp[0][i] = grid[0][i] + tmp[0][i-1]
	}
	for i := 1; i < row; i++ {
		tmp[i][0] = grid[i][0] + tmp[i-1][0]
	}
	for i := 1; i < row; i++ {
		for j := 1; j < cloun; j++ {
			tmp[i][j] = min(tmp[i-1][j], tmp[i][j-1]) + grid[i][j]
		}
	}
	return tmp[row-1][cloun-1]
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
