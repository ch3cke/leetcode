package main

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}
	index := -1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			index = i - 1
			break
		}
	}
	if index == -1 {
		sort.Ints(nums)
		return
	}
	for j := len(nums) - 1; j > index; j-- {
		if nums[j] > nums[index] {
			tmp := nums[j]
			nums[j] = nums[index]
			nums[index] = tmp
			sort.Ints(nums[index+1:])
			fmt.Println(nums)
			return
		}
	}
}
