package main

func transpose(matrix [][]int) [][]int {
	length := len(matrix[0])
	hight := len(matrix)
	result := make([][]int, length)
	for i := 0; i < length; i++ {
		result[i] = make([]int, hight)
	}
	for i := 0; i < hight; i++ {
		for j := 0; j < length; j++ {
			result[j][i] = matrix[i][j]
		}
	}
	return result
}
