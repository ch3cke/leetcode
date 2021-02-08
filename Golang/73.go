package main

func setZeroes(matrix [][]int) {
	rows, clouns := len(matrix), len(matrix[0])
	tmp_table := make([][]int, rows)
	for i := 0; i < rows; i++ {
		tmp_table[i] = make([]int, clouns)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < clouns; j++ {
			tmp_table[i][j] = matrix[i][j]
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < clouns; j++ {
			if tmp_table[i][j] == 0 {
				for a := 0; a < rows; a++ {
					matrix[a][j] = 0
				}
				for b := 0; b < clouns; b++ {
					matrix[i][b] = 0
				}
			}
		}
	}

}
