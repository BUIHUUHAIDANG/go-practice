package main

import "fmt"

type Point struct{
	X int
	Y int
}

type Queue []Point

func (q *Queue) push(p Point){
	*q = append(*q, p)
}
func (q *Queue) pop() Point{
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}

func nearestExit(maze [][]byte, entrance []int) int {
	n := len(maze)
	m := len(maze[0])

	var q Queue

	q.push(Point{entrance[0], entrance[1]})

	maze[entrance[0]][entrance[1]] = '+'

	directions := []Point{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	steps := 0

	for len(q) > 0 {

		size := len(q)

		for i := 0; i < size; i++ {

			cur := q.pop()

			for _, d := range directions {

				nx := cur.X + d.X
				ny := cur.Y + d.Y

				if nx < 0 || ny < 0 || nx >= n || ny >= m {
					continue
				}

				if maze[nx][ny] != '.' {
					continue
				}

				if nx == 0 || ny == 0 || nx == n-1 || ny == m-1 {
					return steps + 1
				}

				maze[nx][ny] = '+'

				q.push(Point{nx, ny})
			}
		}

		steps++
	}

	return -1
}

func main(){
	// maze := [][]byte{
	// {'+', '+', '.', '+'},
	// {'.', '.', '.', '+'},
	// {'+', '+', '+', '.'},
	// }
	// entrance := []int{1,2}
	// maze := [][]byte{
	// 	{'+','+','+'},
	// 	{'.','.','.'},
	// 	{'+','+','+'},
	// }
	// entrance := []int{1,0}
	maze := [][]byte{
		{'.','+'},
	}
	entrance := []int{0,0}
	fmt.Print(nearestExit(maze,entrance))
}