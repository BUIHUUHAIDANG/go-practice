package main

import "fmt"

func bfs(n int ,k int,listAdj [][]int) []bool{
	visited := make([]bool,n)
	q := []int{}
	q = append(q, k)
	visited[k] = true
	for len(q) != 0{
		p := q[0]
		q = q[1:]
		for _,x := range listAdj[p]{
			if visited[x] != true{
				q = append(q, x)
				visited[x] = true
			}
		}
	}
	return visited
}

func remainingMethods(n int, k int, invocations [][]int) []int {
	listAdj := make([][]int,n)
	for _,x  := range invocations{
		listAdj[x[0]] = append(listAdj[x[0]],x[1])
	}
	fmt.Printf("listAdj: %v\n",listAdj)
	suspicious := bfs(n,k,listAdj)
	
	for _, x := range invocations {
		a := x[0]
		b := x[1]

		if !suspicious[a] && suspicious[b] {
			ans := []int{}

			for i := 0; i < n; i++ {
				ans = append(ans, i)
			}

			return ans
		}
	}


	ans := []int{}

	for i := 0; i < n; i++ {
		if !suspicious[i] {
			ans = append(ans, i)
		}
	}

	return ans
}

func main(){
	ans := remainingMethods(4,1,[][]int{{1,2},{0,1},{2,0}})
	fmt.Println(ans)
}