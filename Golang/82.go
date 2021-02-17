package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	h := &ListNode{}
	h.Next = head
	pre := h
	cur := head
	for cur != nil {
		flag := false
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
			flag = true
		}
		if flag == false {
			pre = cur
		} else {
			pre.Next = cur.Next
		}
		cur = cur.Next
	}
	return h.Next
}
