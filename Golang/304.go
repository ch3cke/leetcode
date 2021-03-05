package main

type NumMatrix struct {
	matrix [][]int
}

func Constructor_1(matrix [][]int) NumMatrix {
	result := NumMatrix{
		matrix: matrix,
	}
	return result
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	result := 0
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			result += this.matrix[i][j]
		}
	}
	return result
}
