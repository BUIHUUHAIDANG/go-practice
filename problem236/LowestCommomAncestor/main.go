package main

import "fmt"


type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

func NewNode(val int) *TreeNode{
	return &TreeNode{
		Val: val,
		Left: nil,
		Right: nil,
	}
}



func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    if root == p || root == q {
        return root
    }

    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)

    if left != nil && right != nil {
        return root
    }

    if left != nil {
        return left
    }

    return right
}

func main(){
	root := NewNode(3)
	p1 := NewNode(1)
	p2 := NewNode(2)
	p0 := NewNode(0)
	p4 := NewNode(4)
	p5 := NewNode(5)
	p6 := NewNode(6)
	p7 := NewNode(7)
	p8 := NewNode(8)
	root.Left = p5
	root.Right = p1
	p5.Left = p6
	p5.Right = p2
	p1.Left = p0
	p1.Right =p8
	p2.Left = p7
	p2.Right = p4 
	a := lowestCommonAncestor(root,p5,p4)
	if a == nil{
		fmt.Println("hello")
	}

}