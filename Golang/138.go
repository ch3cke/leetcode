package main

func copyRandomList(head *Node) *Node {
	result := new(Node)
	m := map[*Node]*Node{}
	curr := head
	p := result
	for curr != nil {
		p.Next = &Node{Val: curr.Val}
		m[curr] = p.Next
		curr = curr.Next
		p = p.Next
	}

	curr = head
	p = result.Next
	for curr != nil {
		p.Random = m[curr.Random]
		p = p.Next
		curr = curr.Next
	}
	return result.Next
}
