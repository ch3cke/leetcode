package main

func connect_117(root *Node) *Node {
	tmp_Node := []*Node{}
	if root == nil {
		return nil
	}
	tmp_Node = append(tmp_Node, root)
	for {
		if len(tmp_Node) == 0 {
			break
		}
		tmp := []*Node{}
		for i := 0; i < len(tmp_Node); i++ {
			if i == len(tmp_Node)-1 {
				tmp_Node[i].Next = nil
			} else {
				tmp_Node[i].Next = tmp_Node[i+1]
			}
			if tmp_Node[i].Left != nil {
				tmp = append(tmp, tmp_Node[i].Left)
			}
			if tmp_Node[i].Right != nil {
				tmp = append(tmp, tmp_Node[i].Right)
			}
		}
		tmp_Node = tmp
	}
	return root
}
