package main

import "fmt"

type Node struct {
	Val int
	Next *Node 
}

type ListNode struct {
	head *Node
}

func (l *ListNode) push(num int){
	newNode := &Node{
		Val: num,
	}
	if l.head == nil {
		l.head = newNode
		return
	}
	curr := l.head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
}

func deleteMiddle(l *ListNode) *Node {
	if l.head == nil || l.head.Next == nil {
		return nil 
	}
	slow := l.head
	right := l.head 
	var prev *Node = nil
	for right != nil && right.Next != nil {
		prev = slow
		slow = slow.Next
		right = right.Next.Next
	}
	prev.Next = slow.Next
	return l.head
}

func (l *ListNode) print(){
	tmp := l.head
	for tmp != nil {
		fmt.Printf("%v ->",tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println()
}

func main(){
	node1 := Node{Val: 1}
	node2 := Node{Val: 3}
	node3 := Node{Val: 4}
	node4 := Node{Val: 7}
	node5 := Node{Val: 1}
	node6 := Node{Val: 2}
	node7 := Node{Val: 6}
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	node5.Next = &node6
	node6.Next = &node7
	l := ListNode{}
	l.head = &node1 
	l.print()
	deleteMiddle(&l)
	l.print()
}