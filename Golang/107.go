package main

func levelOrderBottom(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	Node := []*TreeNode{root}
	tmp_result := []int{}
	for {
		if len(Node) == 0 {
			break
		}
		tmp := []*TreeNode{}
		tmp_result = []int{}
		for i := 0; i < len(Node); i++ {
			tmp_result = append(tmp_result, Node[i].Val)
			if Node[i].Left != nil {
				tmp = append(tmp, Node[i].Left)
			}
			if Node[i].Right != nil {
				tmp = append(tmp, Node[i].Right)
			}
		}
		tmp_result_1 := [][]int{}
		tmp_result_1 = append(tmp_result_1, tmp_result)
		result = append(tmp_result_1, result...)
		Node = tmp
	}
	return result
}
