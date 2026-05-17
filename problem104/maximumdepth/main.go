package main

import "fmt"

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left)+1,maxDepth(root.Right)+1)
}

func main(){
	root := TreeNode{}
	node1 := &TreeNode{
		Val: 3,
	}
	node2 := &TreeNode{
		Val: 9,
	}
	node3 := &TreeNode{
		Val: 20,
	}
	node4 := &TreeNode{
		Val: 15,
	}
	node5 := &TreeNode{
		Val: 7,
	}
	root = *node1
	root.Left = node2
	root.Right = node3
	node3.Left = node4
	node3.Right = node5
	fmt.Print(maxDepth(&root))
}