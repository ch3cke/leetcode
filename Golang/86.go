package main

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	Max := &ListNode{}
	p1 := Max
	Min := &ListNode{}
	p2 := Min
	for {
		if head.Val >= x {
			tmp_node := &ListNode{Val: head.Val}
			Max.Next = tmp_node
			Max = Max.Next
		} else {
			tmp_node := &ListNode{Val: head.Val}
			Min.Next = tmp_node
			Min = Min.Next
		}
		if head.Next == nil {
			break
		} else {
			head = head.Next
		}
	}
	Min.Next = p1.Next
	return p2.Next
}
