package main

func sumNumbers(root *TreeNode) int {
	result := 0
	tmp := 0
	var dfs func(root *TreeNode, tmp_num int)
	dfs = func(root *TreeNode, tmp_num int) {
		if (root.Left == nil) && (root.Right == nil) {
			tmp_num = tmp_num*10 + root.Val
			result += tmp_num
			return
		}
		if root.Left != nil {
			dfs(root.Left, root.Val+tmp_num*10)
		}
		if root.Right != nil {
			dfs(root.Right, root.Val+tmp_num*10)
		}
	}
	dfs(root, tmp)
	return result
}
