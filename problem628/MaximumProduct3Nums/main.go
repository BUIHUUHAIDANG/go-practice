package main

import "fmt"

func maximumProduct(nums []int) int {
	num1,num2,num3 := -1000,-1000,-1000
	mn1,mn2 := 1000,1000
	for _,x := range nums{
		if x >= num1 {
			num3 = num2
			num2 = num1
			num1 = x
		}else if x >= num2{
			num3 = num2
			num2 = x 
		}else if x >= num3{
			num3 = x
		}
		// fmt.Printf("Num1: %v\n",num1)
		// fmt.Printf("Num2: %v\n",num2)
		// fmt.Printf("Num3: %v\n",num3)
		if x <= mn1{
			mn2 = mn1
			mn1 = x
		}else if x <= mn2{
			mn2 = x
		}
		// fmt.Printf("mn1: %v\n",mn1)
		// fmt.Printf("mn2: %v\n",mn2)
	}
	return max((num1*num2*num3),(num1*mn1*mn2))
}

func main(){
	ans := maximumProduct([]int{45,3,534,5,-1000,100,-500})
	fmt.Printf("ans: %v\n",ans)
}