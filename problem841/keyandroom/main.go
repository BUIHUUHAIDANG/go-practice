package main

import "fmt"

type Stack []int

func (s *Stack) push(num int){
	*s = append(*s, num)
}

func (s *Stack) pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func canVisitAllRooms(rooms [][]int) bool {
	var s Stack
	for _,v := range rooms[0]{
		s.push(v)
	}
	visited := make([]int,len(rooms))
	visited[0] = 1
	for len(s) > 0 {
		tmp := s.pop()
		visited[tmp] = 1
		for _,v := range rooms[tmp]{
			if visited[v] ==1{
				continue
			}else{
				s.push(v)
			}
		}
	}
	for _,v := range visited{
		if v == 0 {
			return false
		}
	}
	return true
}

func main(){
	matrix := [][]int{{1,3},{3,0,1},{2},{0}}
	fmt.Print(canVisitAllRooms(matrix))
}