package main

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	j := len(nums) - 1
	if j < 0 {
		return 0
	}
	if j == 0 {
		if nums[i] == val {
			return 0
		}
		return 1
	}
	for {
		if nums[i] == val {
			tmp := nums[j]
			nums[j] = val
			nums[i] = tmp
			j = j - 1
		} else {
			i = i + 1
		}
		if i == j+1 {
			break
		}
	}
	fmt.Println(nums)
	return j + 1
}
