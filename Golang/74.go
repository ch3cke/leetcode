package main

func searchMatrix(matrix [][]int, target int) bool {
	if matrix[0][0] > target {
		return false
	}
	if matrix[0][0] == target {
		return true
	}
	rows, cloums := len(matrix), len(matrix[0])
	if rows == 1 {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[0][j] == target {
				return true
			}
		}
		return false
	} else {
		for i := 0; i < rows; i++ {
			if matrix[i][0] == target {
				return true
			}
			if matrix[i][0] > target {
				for j := 0; j < len(matrix[i-1]); j++ {
					if matrix[i-1][j] == target {
						return true
					}
				}
			}
		}
		if matrix[rows-1][cloums-1] >= target {
			for j := 0; j < len(matrix[rows-1]); j++ {
				if matrix[rows-1][j] == target {
					return true
				}
			}
		}
		return false
	}
}
