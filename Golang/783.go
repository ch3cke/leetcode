package main

import (
	"fmt"
	"sort"
)

func minDiffInBST(root *TreeNode) int {
	min_num := []int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		min_num = append(min_num, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	sort.Ints(min_num)
	fmt.Println(min_num)
	result := 1000
	for i := 1; i < len(min_num); i++ {
		if (min_num[i] - min_num[i-1]) < result {
			result = (min_num[i] - min_num[i-1])
		}
	}
	return result
}
