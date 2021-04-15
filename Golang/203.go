package main

func removeElements(head *ListNode, val int) *ListNode {
	result := &ListNode{Next: head}
	pre, mid := result, head
	if head == nil {
		return head
	}
	for mid != nil {
		if mid.Val == val {
			pre.Next = mid.Next
			mid = pre.Next
		} else {
			pre = pre.Next
			mid = mid.Next
		}
	}
	return result.Next
}
