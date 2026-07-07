package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
    for i:= 2; i< len(cost);i++{
		cost[i] = cost[i] + min(cost[i-1],cost[i-2])
	}
	return min(cost[len(cost)-1],cost[len(cost)-2])
}

func main(){
	cost := []int{10,15,20}
	ans := minCostClimbingStairs(cost)
	fmt.Print(ans)

}