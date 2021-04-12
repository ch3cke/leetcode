package main

import "sort"

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	result := 0
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		s := nums[i] - nums[i-1]
		if s > result {
			s = result
		}
	}
	return result
}
