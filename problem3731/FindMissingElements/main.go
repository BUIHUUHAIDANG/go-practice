package main

import "fmt"

func findMissingElements(nums []int) []int {
	listHundred := make([]bool,100)
	ans := []int{}
	MaxNum := -1
	MinNum := 101
	for _,x := range nums{
		MaxNum = max(MaxNum,x)
		MinNum = min(MinNum,x)
		listHundred[x] = true
	}
	if MaxNum - MinNum + 1 == len(nums){
		return []int{}
	}
	for i := MinNum ; i <= MaxNum; i++{
		if listHundred[i] != true{
			ans = append(ans, i)
		}
	}
	return ans	
}

func main(){
	ans := findMissingElements([]int{5,1})
	fmt.Printf("ans: %v\n",ans)
}