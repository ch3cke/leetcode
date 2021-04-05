package main

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	tmp := []*ListNode{head}
	end := head.Next
	if end == nil {
		return
	}
	for end != nil {
		tmp = append(tmp, end)
		end = end.Next
	}
	i := 0
	j := len(tmp) - 1
	for i < j {
		tmp[i].Next = tmp[j]
		i += 1
		tmp[j].Next = tmp[i]
		j -= 1
	}
}
