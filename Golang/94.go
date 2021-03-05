package main

import "fmt"

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	var getNode func(Node *TreeNode)
	getNode = func(Node *TreeNode) {
		if Node == nil {
			return
		}
		getNode(Node.Left)
		result = append(result, Node.Val)
		getNode(Node.Right)
	}
	getNode(root)
	fmt.Println(result)
	return result
}
