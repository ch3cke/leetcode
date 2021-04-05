package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	tmp := map[*ListNode]int{}
	for headA != nil {
		tmp[headA] = headA.Val
		headA = headA.Next
	}
	for headB != nil {
		_, ok := tmp[headB]
		if ok {
			return headB
		} else {
			headB = headB.Next
		}
	}
	return nil
}
