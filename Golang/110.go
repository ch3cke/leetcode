package main

func isBalanced(root *TreeNode) bool {
	return checkTree(root) >= 0
}

func checkTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left_hight := checkTree(root.Left)
	right_hight := checkTree(root.Right)
	if left_hight == -1 || right_hight == -1 || abs(left_hight-right_hight) > 1 {
		return -1
	}
	return max(left_hight, right_hight) + 1
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return 0 - a
	}
	return a
}
