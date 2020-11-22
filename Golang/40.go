package main

import (
	"reflect"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	comb := []int{}
	var ans [][]int
	var dfs func(target int, idx int)
	dfs = func(target, idx int) {
		if target == 0 {
			tmp := append([]int(nil), comb...)
			sort.Ints(tmp)
			ans = append(ans, tmp)
			return
		}
		if idx == len(candidates) {
			return
		}
		dfs(target, idx+1)
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx+1)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return killRepetion(ans)
}

func killRepetion(nums [][]int) [][]int {
	newRes := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		flag := false
		for j := i + 1; j < len(nums); j++ {
			if reflect.DeepEqual(nums[i], nums[j]) {
				flag = true
				break
			}
		}
		if !flag {
			newRes = append(newRes, nums[i])
		}
	}
	return newRes
}
