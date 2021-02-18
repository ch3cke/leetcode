package main

import (
	"reflect"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	result := [][]int{}
	length := 1 << len(nums)
	for m := 0; m < length; m++ {
		tmp := []int{}
		for i := 0; i < len(nums); i++ {
			if (m & (1 << i)) != 0 {
				tmp = append(tmp, nums[i])
			}
		}
		sort.Ints(tmp)
		if !check(result, tmp) {
			result = append(result, tmp)
		}
	}
	return result
}

func check(flag [][]int, test []int) bool {
	for i := 0; i < len(flag); i++ {
		if reflect.DeepEqual(flag[i], test) {
			return true
		}
	}
	return false
}
