package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	flag := true
	var getNode func(s *TreeNode, d *TreeNode)
	getNode = func(s *TreeNode, d *TreeNode) {
		if s == nil || d == nil {
			if (s == nil && d != nil) || (s != nil && d == nil) {
				flag = false
			}
			return
		}
		getNode(s.Left, d.Left)
		if s.Val != d.Val {
			flag = false
		}
		getNode(s.Right, d.Right)
	}
	getNode(p, q)
	return flag
}
