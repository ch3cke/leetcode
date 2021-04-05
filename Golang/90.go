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

func subsetsWithDup_1(nums []int) [][]int {
	result := [][]int{}
	if len(nums) == 0 {
		return result
	}
	sort.Ints(nums)
	tmp := []int{}
	var dfs func(bool, int)
	dfs = func(choose bool, cur int) {
		if cur == len(nums) {
			result = append(result, append([]int(nil), tmp...))
			return
		}
		dfs(false, cur+1)
		if !choose && cur > 0 && nums[cur-1] == nums[cur] {
			return
		}
		tmp = append(tmp, nums[cur])
		dfs(true, cur+1)
		tmp = tmp[:len(tmp)-1]
	}
	dfs(false, 0)
	return result
}
