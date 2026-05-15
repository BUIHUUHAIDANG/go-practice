package main

import "fmt"

func pivotIndex(nums []int) int {
    a := make([]int,len(nums))
	for i:= 0;i<len(nums);i++{
		if i == 0{
			a[0] = nums[0]
		}else{
			a[i] = a[i-1]+nums[i]
		}
	}
	for i:= 0; i< len(nums);i++{
		var leftsum int 
		var rightsum int
		if i == 0 {
			leftsum = 0
			rightsum = a[len(nums)-1]-a[i]
		}else if i == len(nums)-1{
			leftsum = a[i-1]
			rightsum = 0
		}else{
			leftsum = a[i-1]
			rightsum = a[len(nums)-1]-a[i]

		}

		if leftsum == rightsum {
			return i
		}
	}
	return -1
}

func main(){
	fmt.Print(pivotIndex([]int{2,1,-1}))	
}