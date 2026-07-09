package main

import "fmt"


// func dfs(matrix [][]int,s int,e int) bool{
// 	visited := make([]bool,len(matrix[0]))
// 	stack := []int{}
// 	stack = append(stack, s)
// 	visited[s] = true
// 	for len(stack) != 0 {
// 		u := stack[len(stack)-1]
// 		if u == e {
// 			return true
// 		}
// 		stack = stack[:len(stack)-1]
// 		for i := 0; i < len(matrix[0]);i++{
// 			if visited[i] == false && matrix[u][i] == 1{
// 				stack = append(stack, i)
// 				visited[i] = true
// 			}
// 		}
// 	}
// 	return false

// }

// func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
// 	matrix := make([][]int,n)
// 	for i := range matrix {
// 		matrix[i] = make([]int, n)
// 	}
// 	for i:= 0; i<len(nums); i++{
// 		matrix[i][i] = 1
// 		if i + 1 < len(nums) && max(nums[i]-nums[i+1],nums[i+1]-nums[i]) <= maxDiff{
// 			matrix[i][i+1] = 1
// 			matrix[i+1][i] = 1
// 		} 
// 	} 
// 	ans := make([]bool,len(queries))
// 	for i := 0;i<len(queries);i++{
// 		if queries[i][0] == queries[i][1]{
// 			ans[i] = true
// 		}else{
// 			ans[i] = dfs(matrix,queries[i][0],queries[i][1])
// 		}
// 	}
// 	return ans
// }

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
    comp := make([]int, n)

    id := 0
    comp[0] = id
	// set id a chain have a abs(nums[i]-nums[i-1]) below maxdiff have same id  
    for i := 1; i < n; i++ {
        if abs(nums[i]-nums[i-1]) > maxDiff {
            id++
        }
        comp[i] = id
    }

    ans := make([]bool, len(queries))
    for i, q := range queries {
		// if same id is true else false 
        ans[i] = comp[q[0]] == comp[q[1]]
    }

    return ans
}


func main(){
	//Input: n = 2, nums = [1,3], maxDiff = 1, queries = [[0,0],[0,1]]
	//n = 4, nums = [2,5,6,8], maxDiff = 2, queries = [[0,1],[0,2],[1,3],[2,3]]
	var ans []bool = pathExistenceQueries(4, []int{2,5,6,8},2,[][]int{{0,1},{0,2},{1,3},{2,3}})
	fmt.Printf("ANS: %v\n",ans)
}