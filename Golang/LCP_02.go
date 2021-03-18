package main

func fraction(cont []int) []int {
	if len(cont) == 1 {
		return append(cont, 1)
	}
	result := make([]int, 2)
	result[0] = 1
	result[1] = cont[len(cont)-1]
	for i := len(cont) - 2; i >= 0; i-- {
		result[0] += cont[i] * result[1]
		tmp := result[1]
		result[1] = result[0]
		result[0] = tmp
	}
	tmp := result[1]
	result[1] = result[0]
	result[0] = tmp
	return result
}
