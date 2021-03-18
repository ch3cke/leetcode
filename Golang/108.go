package main

func sortedArrayToBST(nums []int) *TreeNode {
	return tmp_func(nums, 0, len(nums)-1)
}

func tmp_func(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = tmp_func(nums, left, mid-1)
	root.Right = tmp_func(nums, mid+1, right)
	return root
}
