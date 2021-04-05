package main

import (
	"fmt"
	"sort"
)

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tmp := []int{}
	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}
	fmt.Println(tmp)
	sort.Ints(tmp)
	return init_List(tmp)
}
