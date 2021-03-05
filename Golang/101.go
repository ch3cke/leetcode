package main

func isSymmetric(root *TreeNode) bool {
	return check_1(root, root)
}

func check_1(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check_1(p.Left, q.Right) && check_1(p.Right, q.Left)
}
