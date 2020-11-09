package main

//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]
//
//
// 示例 2：
//
//
//输入：head = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 100] 内
// 0 <= Node.val <= 100
//
// Related Topics 链表
// 👍 733 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	index := 1
	grad := &ListNode{}
	grad = head
	for {
		if head == nil {
			break
		}
		if head.Next == nil {
			break
		}
		tmpNode1 := head
		tmpNode2 := head.Next
		tmp := tmpNode1.Val
		tmpNode1.Val = tmpNode2.Val
		tmpNode2.Val = tmp
		index += 2
		head = head.Next
		head = head.Next
	}

	return grad
}

//leetcode submit region end(Prohibit modification and deletion)
