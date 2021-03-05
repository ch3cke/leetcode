package main

import "fmt"

func countBits(num int) []int {
	result := []int{0}
	if num == 0 {
		return result
	}
	result = append(result, 1)
	if num == 1 {
		return result
	}
	for {
		tmp := result
		for i := 0; i < len(result); i++ {
			tmp = append(tmp, result[i]+1)
		}
		fmt.Println(tmp)
		result = tmp
		if len(result) > num {
			break
		}
	}
	return result[:num+1]
}
