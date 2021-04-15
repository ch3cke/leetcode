package main

func countNodes(root *TreeNode) int {
	result := 0
	if root == nil {
		return result
	}
	tmp := []*TreeNode{root}
	for len(tmp) != 0 {
		result += len(tmp)
		tmp_list := []*TreeNode{}
		for i := 0; i < len(tmp); i++ {
			if tmp[i].Left != nil {
				tmp_list = append(tmp_list, tmp[i].Left)
			}
			if tmp[i].Right != nil {
				tmp_list = append(tmp_list, tmp[i].Right)
			}
		}
		tmp = tmp_list
	}
	return result
}
