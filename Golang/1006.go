package main

import "fmt"

func clumsy(N int) int {
	result := 0
	if N > 3 {
		result = N*(N-1)/(N-2) + (N - 3)
	} else if N == 3 {
		return N * (N - 1) / (N - 2)
	} else if N == 2 {
		return 2
	} else if N == 1 {
		return 1
	} else {
		return 0
	}
	i := N - 4
	for {
		if i < 4 {
			break
		}
		tmp := i * (i - 1) / (i - 2)
		result = result - tmp
		result += (i - 3)
		i = i - 4
	}
	fmt.Println(result)
	fmt.Println(i)
	if i == 3 {
		tmp := i * (i - 1) / (i - 2)
		result -= tmp
	} else if i == 2 {
		tmp := i * (i - 1)
		result -= tmp
	} else if i == 1 {
		tmp := 1
		result -= tmp
	}
	return result
}
