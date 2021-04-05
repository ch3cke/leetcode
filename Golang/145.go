package main

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		result = append(result, root.Val)
	}
	dfs(root)
	return result
}
