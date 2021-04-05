package main

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		result = append(result, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return result
}
