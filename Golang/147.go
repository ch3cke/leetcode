package main

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	result := &ListNode{}
	tmp := result.Next
	for {
		if head == nil {
			break
		}
		to_insert := head
		head = head.Next
		to_insert.Next = nil
		if result.Next == nil {
			result.Next = to_insert
			continue
		} else {
			tmp = result.Next
		}
		if to_insert.Val <= tmp.Val {
			to_insert.Next = tmp
			result.Next = to_insert
		} else {
			for {
				if tmp.Next == nil || (tmp.Val <= to_insert.Val && tmp.Next.Val >= to_insert.Val) {
					to_insert.Next = tmp.Next
					tmp.Next = to_insert
					break
				} else {
					tmp = tmp.Next
				}
			}
		}
	}
	return result.Next
}
