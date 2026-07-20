package main

import "fmt"


// func shiftGrid(grid [][]int, k int) [][]int {
// 	for k != 0{
// 		for i:= 0; i<len(grid);i++{
// 			for j:= 0;j<len(grid[0]);j++{
// 				grid[0][0],grid[i][j] = grid[i][j],grid[0][0]
// 			}
// 		}
// 		k--
// 	}
// 	return  grid
// }

func shiftGrid(grid [][]int, k int) [][]int {
	x,y := len(grid),len(grid[0])
	NewGrid := make([][]int,x)
	for i := range NewGrid{
		NewGrid[i] = make([]int, y)
	}
	for i := 0 ;i <x;i++{
		for j := 0; j <y;j++{
			idx := i*y +j
			Newidx := (idx+k)%(x*y)
			NewGrid[Newidx/y][Newidx%y] = grid[i][j]
		}
	}
	return NewGrid	
}

func main(){
	// grid := [][]int{
	// 	{3,8,1,9},
	// 	{19,7,2,5},
	// 	{4,6,11,10},
	// 	{12,0,21,13},
	// }
	grid := [][]int{
		{1},
		{2},
		{3},
		{4},
		{7},
		{6},
		{5},
	}
	ans := shiftGrid(grid,23)
	for _,x := range ans{
		fmt.Printf("row : %v\n",x)
	}
}