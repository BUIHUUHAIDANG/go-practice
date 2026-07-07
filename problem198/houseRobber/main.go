package main

import "fmt"

func maxArray(nums []int, x int) int {
	maxnum := 0
	for i:= 0;i <x;i++{
		if nums[i]>maxnum{
			maxnum =nums[i]
		}
	} 
	return maxnum
}

func rob(nums []int) int {
	if len(nums) == 1{
        return nums[0]
    }
    for i := 2 ;i<len(nums);i++{
		nums[i] = nums[i]+ maxArray(nums,i-1)
	}
	return  max(nums[len(nums)-1],nums[len(nums)-2])
}

func main(){
	house := []int{1,2,3,1}
	ans := rob(house)
	fmt.Print(ans)
}