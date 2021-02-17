package main

func deleteDuplicates_1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p1 := head
	p2 := head.Next

	for {
		if p2 == nil {
			break
		}
		if p1.Val == p2.Val {
			p1.Next = p2.Next
			p2 = p2.Next
		} else {
			p1 = p1.Next
			p2 = p2.Next
		}
	}
	return head
}
