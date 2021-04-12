package main

func rightSideView(root *TreeNode) []int {
	tmp_list := []*TreeNode{root}
	result := []int{}
	for len(tmp_list) != 0 {
		tmp := []*TreeNode{}
		for i := 0; i < len(tmp_list); i++ {
			if i == 0 {
				result = append(result, tmp_list[i].Val)
			}
			if tmp_list[i].Right != nil {
				tmp = append(tmp, tmp_list[i].Right)
			}
			if tmp_list[i].Left != nil {
				tmp = append(tmp, tmp_list[i].Left)
			}
		}
		tmp_list = tmp
	}
	return result
}
