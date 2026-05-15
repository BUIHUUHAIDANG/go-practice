package main

import "fmt"

func moveZeroes(nums []int)  {
	i:= 0
	j:= 0
	for j < len(nums){
		for i<j &&nums[i] != 0 {
			i++
		}
		if nums[j] != 0 && nums[i]== 0 {
			nums[i],nums[j] = nums[j],nums[i]
		}
		j++  
	}
}

func main(){
	nums:= []int{1}
	moveZeroes(nums)
	fmt.Print(nums)
}