package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func swap(l1 *ListNode, l2 *ListNode) {
	tmp := l1.Val
	l2.Val = tmp
	l1.Val = tmp
}

func initList(l *ListNode, table []int) {
	l2 := l
	l2.Val = table[0]
	for i := 1; i < len(table); i++ {
		l3 := ListNode{Val: table[i]}
		l2.Next = &l3
		l2 = &l3
	}
}

func showList(l *ListNode) {
	for l2 := l; l2 != nil; l2 = l2.Next {
		fmt.Println(l2.Val, "\t")
	}
}

func main() {
	//[-10,-10,-9,-4,1,6,6]
	//[-7]
	l := &ListNode{}
	// l2 := []int{1,2,3,4}
	// initList(l, l2)
	showList(swapPairs(l))
}
