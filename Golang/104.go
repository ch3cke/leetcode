package main

func maxDepth(root *TreeNode) int {
	result := 0
	if root == nil {
		return result
	}
	var dfs func(q *TreeNode, s int)
	s := 0
	dfs = func(q *TreeNode, s int) {
		if q == nil {
			if result < s {
				result = s
			}
			return
		}
		dfs(q.Left, s+1)
		dfs(q.Right, s+1)
	}
	dfs(root, s)
	return result
}
