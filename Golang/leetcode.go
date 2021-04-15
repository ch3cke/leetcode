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

//type Node struct {
//	Val   int
//	Left  *Node
//	Right *Node
//	Next  *Node
//}

//type Node struct {
//	Val       int
//	Neighbors []*Node
//}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
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

func countPrimes(n int) int {
	signs := make([]bool, n)
	for i := 0; i < n; i++ {
		signs[i] = true
	}
	result := 0
	for i := 2; i < n; i++ {
		if signs[i] {
			result += 1
			for j := i + i; j < n; j += i {
				signs[j] = false
			}
		}
	}
	return result
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	tmp_map := map[uint8]uint8{}
	tmp_map2 := map[uint8]uint8{}
	for i := 0; i < len(s); i++ {
		i1, ok := tmp_map[s[i]]
		i2, ok2 := tmp_map2[t[i]]
		if ok {
			if i1 != t[i] {
				return false
			}
		} else if ok2 {
			if i2 != s[i] {
				return false
			}
		} else {
			tmp_map[s[i]] = t[i]
			tmp_map2[t[i]] = s[i]
		}
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	result := &ListNode{}
	for head != nil {
		result.Val = head.Val
		tmp1 := &ListNode{
			Next: result,
		}
		result = tmp1
		head = head.Next
	}
	return result.Next
}

func rob_2(nums []int) int {
	var tmp func([]int) int
	if len(nums) == 1 {
		return nums[0]
	}
	tmp = func(ints []int) int {
		cur, pre := 0, 0
		for i := 0; i < len(ints); i++ {
			cur, pre = max_num2(pre+ints[i], cur), cur
		}
		return cur
	}
	return max_num2(tmp(nums[:len(nums)-1]), tmp(nums[1:]))
}

func max_num2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func combinationSum3(k int, n int) [][]int {
	tmp := []int{}
	result := [][]int{}
	var dfs func(rest, cur int)
	dfs = func(rest, cur int) {
		if len(tmp) == k && rest == 0 {
			result = append(result, append([]int{}, tmp...))
			return
		}
		if len(tmp)+10-cur < k || rest < 0 {
			return
		}
		dfs(rest, cur+1)
		tmp = append(tmp, cur)
		dfs(rest-cur, cur+1)
		tmp = tmp[:len(tmp)-1]
	}
	dfs(n, 1)
	return result
}

func containsDuplicate(nums []int) bool {
	tmp := map[int]bool{}
	for _, i2 := range nums {
		_, ok := tmp[i2]
		if ok {
			return true
		} else {
			tmp[i2] = false
		}
	}
	return false
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	tmp := (getDistance(A, C) * getDistance(B, D)) + (getDistance(E, G) * getDistance(F, H))
	if a(A, B, C, D, E, F, G, H) {
		return tmp - ((C - E) * (H - B))
	} else {
		return tmp
	}
}

func a(A int, B int, C int, D int, E int, F int, G int, H int) bool {
	if E > A && E < C && F > B && F < D {
		return true
	}
	if G > A && G < C && H > B && H < D {
		return true
	}
	return false
}

func getDistance(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}
