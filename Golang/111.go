package main

func minDepth(root *TreeNode) int {
	result := 0
	if root == nil {
		return result
	}
	tmp := []*TreeNode{root}
	for {
		if len(tmp) == 0 {
			return result
		}
		tmp_Node := []*TreeNode{}
		result += 1
		for i := 0; i < len(tmp); i++ {
			if (tmp[i].Right == nil) && (tmp[i].Left == nil) {
				return result
			} else {
				if tmp[i].Left != nil {
					tmp_Node = append(tmp_Node, tmp[i].Left)
				}
				if tmp[i].Right != nil {
					tmp_Node = append(tmp_Node, tmp[i].Right)
				}
			}
		}
		tmp = tmp_Node
	}
}
