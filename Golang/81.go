package main

import "sort"

func search_1(nums []int, target int) bool {
	sort.Ints(nums)
	index := 0
	end := len(nums) - 1
	for {
		if index == end {
			if nums[index] == target {
				return true
			} else {
				return false
			}
		}
		if end-index == 1 {
			if nums[index] == target || nums[end] == target {
				return true
			} else {
				return false
			}
		}
		mid := (index + end) / 2
		if nums[mid] > target {
			end = mid
		} else {
			index = mid
		}
	}
}
