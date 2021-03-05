package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	pre := &ListNode{}
	pre.Next = head
	mid := head
	head = pre
	if head.Next == nil {
		return head
	}
	for i := 1; i < left; i++ {
		pre = pre.Next
		mid = mid.Next
	}
	tmp := &ListNode{}
	s := tmp
	for i := left; i <= right; i++ {
		d := &ListNode{
			Val: mid.Val,
		}
		d.Next = s
		s = d
		mid = mid.Next
	}
	pre.Next = s
	tmp = s
	for {
		if tmp.Next.Next == nil {
			break
		} else {
			tmp = tmp.Next
		}
	}
	tmp.Next = mid
	return head.Next
}
