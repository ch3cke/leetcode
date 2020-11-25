package main

import (
	"fmt"
	"reflect"
	"sort"
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
	fmt.Println(permute([]int{0, -1, 1, 1}))
	fmt.Println(reflect.DeepEqual([]int{1, 2, 3}, []int{1, 2, 3}))
}
