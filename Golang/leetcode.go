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

func init_List(nums []int) *ListNode {
	result := &ListNode{}
	index := result
	index.Val = nums[0]
	for i := 1; i < len(nums); i++ {
		tmp := &ListNode{
			Val:  nums[i],
			Next: nil,
		}
		index.Next = tmp
		index = tmp
	}
	index.Next = nil
	return result
}

func main() {
	fmt.Println(transpose([][]int{
		{1, 2, 3},
		{4, 5, 6},
	}))
}
