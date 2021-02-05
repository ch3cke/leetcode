package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 1 {
		return matrix[0]
	}
	if len(matrix[0]) == 1 {
		tmp := []int{}
		for i := 0; i < len(matrix); i++ {
			tmp = append(tmp, matrix[i]...)
		}
		return tmp
	}
	result := []int{}
	result = append(result, count(matrix)...)
	tmp := matrix
	tmp = reco(matrix, len(tmp[0]), len(tmp))
	if len(tmp) != 0 {
		if len(tmp[0]) == 1 {
			for i := 0; i < len(tmp); i++ {
				result = append(result, tmp[i]...)
			}
			return result
		}
	}
	for {
		if len(tmp) == 1 {
			result = append(result, tmp[0]...)
			break
		}
		if len(tmp) == 0 {
			break
		}
		result = append(result, count(tmp)...)
		tmp = reco(tmp, len(tmp[0]), len(tmp))
		if len(tmp) != 0 {
			if len(tmp[0]) == 1 {
				for i := 0; i < len(tmp); i++ {
					result = append(result, tmp[i]...)
				}
				return result
			}
		}
	}
	return result
}

func count(matrix [][]int) []int {
	hight := len(matrix)
	lenth := len(matrix[0])
	h := lenth - 1
	l := 1
	result := []int{}
	result = append(result, matrix[0]...)
	for {
		if l == hight-1 {
			break
		} else {
			result = append(result, matrix[l][h])
		}
		l++
	}
	for {
		if h == 0 {
			break
		} else {
			result = append(result, matrix[l][h])
		}
		h = h - 1
	}
	for {
		if l == 0 {
			break
		} else {
			result = append(result, matrix[l][h])
		}
		l = l - 1
	}

	return result
}

func reco(matrix [][]int, raw int, len int) [][]int {
	tmp := [][]int{}
	for i := 1; i < len-1; i++ {
		tmp_raw := []int{}
		for j := 1; j < raw-1; j++ {
			tmp_raw = append(tmp_raw, matrix[i][j])
		}
		if cap(tmp_raw) != 0 {
			tmp = append(tmp, tmp_raw)
		}
	}
	return tmp
}
