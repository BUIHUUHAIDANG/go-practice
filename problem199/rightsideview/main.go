package main


type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Queue []*TreeNode

func (q *Queue) push(node *TreeNode) {
	*q = append(*q, node)
}
func (q *Queue) pop() *TreeNode{
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}

func (q *Queue) top() *TreeNode{
	return (*q)[0]
}
func rightSideView(root *TreeNode) []int {
	if root == nil{
		return []int{}
	}
	res := make([]int,0)
	var q Queue
	q.push(root)
	for len(q)>0{
		size := len(q)
		for i := 0 ;i <size; i++{
			tmp := q.pop()
			if tmp.Left != nil{
				q.push(tmp.Left)
			}
			if tmp.Right != nil{
				q.push(tmp.Right)
			}
			if i == size -1{
				res = append(res, tmp.Val)
			}

		}
	}
	return res
}

func main(){
	

}