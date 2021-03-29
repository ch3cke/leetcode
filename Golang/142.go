package main

func detectCycle(head *ListNode) *ListNode {
	ListSet := make(map[*ListNode]bool)
	for {
		if head == nil {
			return nil
		}
		_, ok := ListSet[head]
		if ok {
			return head
		} else {
			ListSet[head] = true
		}
		head = head.Next
	}
}
