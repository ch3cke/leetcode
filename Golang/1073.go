package main

func addNegabinary(arr1 []int, arr2 []int) []int {
	result := []int{}
	if len(arr1) < len(arr2) {
		return addNegabinary(arr2, arr1)
	}
	if len(arr1) != len(arr2) {
		tmp := make([]int, len(arr1)-len(arr2))
		arr2 = append(tmp, arr2...)
	}
	length := len(arr2)
	tmp := 0
	for i := length - 1; i >= 0; i-- {
		s := arr1[i] + arr2[i] + tmp
		if s == 2 {
			tmp = -1
			tmp_result := []int{0}
			result = append(tmp_result, result...)
		} else if s == -1 {
			tmp = 1
			tmp_result := []int{1}
			result = append(tmp_result, result...)
		} else if s == 3 {
			tmp = -1
			tmp_result := []int{1}
			result = append(tmp_result, result...)
		} else {
			tmp = 0
			tmp_result := []int{s}
			result = append(tmp_result, result...)
		}
	}
	if tmp == 1 {
		tmp_result := []int{tmp}
		result = append(tmp_result, result...)
	} else if tmp == -1 {
		tmp_result := []int{1, 1}
		result = append(tmp_result, result...)
	}
	for i := 0; i < len(result)-1; i++ {
		if result[i] != 0 {
			break
		} else {
			result = append(result[:i], result[i+1:]...)
			i = i - 1
		}
	}
	return result
}
