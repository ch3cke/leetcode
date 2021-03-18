package main

func sortedListToBST(head *ListNode) *TreeNode {
	tmp_list := getList(head)
	return tmp_func1(tmp_list, 0, len(tmp_list)-1)
}
func tmp_func1(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = tmp_func1(nums, left, mid-1)
	root.Right = tmp_func1(nums, mid+1, right)
	return root
}
func getList(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
