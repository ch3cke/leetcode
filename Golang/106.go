package main

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 1 {
		return &TreeNode{Val: inorder[0]}
	}
	var dfs func(root *TreeNode, left []int)
	dfs = func(root *TreeNode, left []int) {
		if len(left) == 1 {
			root.Val = left[0]
			postorder = postorder[:len(postorder)-1]
			return
		}
		if len(left) == 0 {
			root = nil
			return
		}
		if len(postorder) == 0 {
			return
		}
		root_val := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		local := 0
		for i := 0; i < len(left); i++ {
			if left[i] == root_val {
				local = i
				break
			}
		}
		root.Val = root_val
		root.Left = &TreeNode{}
		root.Right = &TreeNode{}
		if local == len(left)-1 {
			root.Right = nil
		} else {
			dfs(root.Right, left[local+1:])
		}
		if local == 0 {
			root.Left = nil
		} else {
			dfs(root.Left, left[:local])
		}
	}
	result := &TreeNode{}
	dfs(result, inorder)
	return result
}
