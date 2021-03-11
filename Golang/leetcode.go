package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

func calculate_1(s string) int {
	ops := []string{}
	nums := []int{}
	sign := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			tmp := 0
			for {
				if s[i] < '0' || s[i] > '9' {
					nums = append(nums, tmp)
					break
				}
				tmp = tmp*10 + int(s[i]-'0')
				i = i + 1
			}
		} else {
			if s[i] == '*' || s[i] == '/' {
				ops = append(ops, string(s[i]))
			} else {

			}
		}
	}
}

func main() {
	//nums :=[][]int{
	//	{5,4},
	//	{6,4},
	//	{6,7},
	//	{2,3},
	//}
	fmt.Println(calculate("1+1"))
}
