package main

import "fmt"

func pathSum(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	var dfs func(root *TreeNode, nums int, path []int)
	dfs = func(root *TreeNode, nums int, path []int) {
		if root == nil {
			return
		}
		if (root.Left == nil) && (root.Right == nil) && (nums-root.Val == 0) {
			path = append(path, root.Val)
			tmp_path := []int{}
			tmp_path = append(tmp_path, path...)
			result = append(result, tmp_path)
			fmt.Println(result)
			return
		}
		dfs(root.Left, nums-root.Val, append(path, root.Val))
		dfs(root.Right, nums-root.Val, append(path, root.Val))
	}
	tmp := []int{}
	dfs(root, targetSum, tmp)
	return result
}
