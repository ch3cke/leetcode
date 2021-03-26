package main

import "math"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	result[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		result[i][0] = result[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			result[i][j] = minNum(result[i-1][j-1], result[i-1][j]) + triangle[i][j]
		}
		result[i][i] = result[i-1][i-1] + triangle[i][i]
	}
	t := math.MaxInt32
	for i := 0; i < n; i++ {
		t = minNum(t, result[n-1][i])
	}
	return t
}

func minNum(a, b int) int {
	if a > b {
		return b
	}
	return a
}
