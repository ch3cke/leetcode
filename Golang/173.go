package main

type BSTIterator struct {
	nums []int
}

func Constructor_6(root *TreeNode) BSTIterator {
	result := BSTIterator{}
	result.nums = []int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		result.nums = append(result.nums, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return result
}

func (this *BSTIterator) Next() int {
	val := this.nums[0]
	this.nums = this.nums[1:]
	return val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.nums) > 0
}
