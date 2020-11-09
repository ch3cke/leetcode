package main

//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//
//
// 示例：
//
// 输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
//
// Related Topics 链表
// 👍 1363 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	gradu := &ListNode{}
	head := gradu
	for {
		if l1 == nil {
			head.Next = l2
			break
		}
		if l2 == nil {
			head.Next = l1
			break
		}
		if l1.Val < l2.Val {
			head.Next = &ListNode{
				Val: l1.Val,
			}
			l1 = l1.Next
		} else {
			head.Next = &ListNode{
				Val: l2.Val,
			}
			l2 = l2.Next
		}
		head = head.Next
	}
	return gradu.Next
}

//leetcode submit region end(Prohibit modification and deletion)
