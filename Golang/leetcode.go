package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func showList(l *ListNode) {
	for l2 := l; l2 != nil; l2 = l2.Next {
		fmt.Println(l2.Val, "\t")
	}
}
func main() {
	fmt.Println(multiply("401716832807512840963", "167141802233061013023557397451289113296441069"))
}
