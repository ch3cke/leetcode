package main

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return getTree(1, n)
}

func getTree(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	allTrees := []*TreeNode{}
	for i := start; i <= end; i++ {
		leftTree := getTree(start, i-1)
		rightTree := getTree(i+1, end)
		for _, left := range leftTree {
			for _, right := range rightTree {
				currTree := &TreeNode{i, nil, nil}
				currTree.Left = left
				currTree.Right = right
				allTrees = append(allTrees, currTree)
			}
		}
	}
	return allTrees
}
