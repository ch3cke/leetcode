package main

func grayCode(n int) []int {
	result := []int{}
	if n == 0 {
		return append(result, 0)
	}
	result = append(result, 0)
	result = append(result, 1)
	for j := 1; j < n; j++ {
		tmp := []int{}
		for i := 0; i < len(result); i++ {
			tmp = append(tmp, result[i])
		}
		for i := len(result) - 1; i >= 0; i-- {
			tmp = append(tmp, result[i]+(1<<j))
		}
		result = tmp
	}
	return result
}
