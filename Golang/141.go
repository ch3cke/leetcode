package main

func hasCycle(head *ListNode) bool {
	tap := make(map[*ListNode]bool)
	for {
		if head == nil {
			return false
		}
		_, ok := tap[head]
		if !ok {
			tap[head] = true
			head = head.Next
		} else {
			return true
		}
	}
}
