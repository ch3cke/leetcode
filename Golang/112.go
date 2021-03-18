package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	result := false
	var dfs func(root *TreeNode, nums int)
	dfs = func(root *TreeNode, nums int) {
		if root == nil {
			return
		}
		if (nums-root.Val == 0) && (root.Left == nil) && (root.Right == nil) {
			result = true
		}
		dfs(root.Left, nums-root.Val)
		dfs(root.Right, nums-root.Val)
	}
	dfs(root, targetSum)
	return result
}
