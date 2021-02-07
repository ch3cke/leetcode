package main

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	count := 0
	index := head
	index2 := head
	for {
		count++
		if index.Next != nil {
			index = index.Next
		} else {
			break
		}
	}
	index.Next = index2
	step := count - (k % count)
	for i := step; i > 0; i-- {
		index = index.Next
		index2 = index2.Next
	}
	index.Next = nil
	return index2
}
