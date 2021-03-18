package main

func numDistinct(s string, t string) int {
	len_s := len(s)
	len_t := len(t)
	if len_s < len_t {
		return 0
	}
	result := make([][]int, len_s+1)
	for i := 0; i < len_s; i++ {
		result[i] = make([]int, len_t+1)
		result[i][len_t] = 1
	}
	for i := len_s - 1; i >= 0; i-- {
		for j := len_t; j >= 0; j-- {
			if s[i] == t[j] {
				result[i][j] = result[i+1][j+1] + result[i+1][j]
			} else {
				result[i][j] = result[i+1][j]
			}
		}
	}
	return result[0][0]
}
