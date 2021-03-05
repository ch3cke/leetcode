package main

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	count := 1
	if root == nil {
		return result
	}
	tmp_result := []int{}
	tmp_Node := []*TreeNode{root}
	for {
		if len(tmp_Node) == 0 {
			break
		}
		tmp_result = []int{}
		tmp := []*TreeNode{}
		for i := 0; i < len(tmp_Node); i++ {
			tmp_result = append(tmp_result, tmp_Node[i].Val)
			if tmp_Node[i].Right != nil {
				tmp = append(tmp, tmp_Node[i].Right)
			}
			if tmp_Node[i].Left != nil {
				tmp = append(tmp, tmp_Node[i].Left)
			}
		}
		if count%2 == 1 {
			for i, n := 0, len(tmp_result); i < n/2; i++ {
				tmp_result[i], tmp_result[n-1-i] = tmp_result[n-1-i], tmp_result[i]
			}
		}
		result = append(result, tmp_result)
		tmp_Node = tmp
		count += 1
	}
	return result
}
