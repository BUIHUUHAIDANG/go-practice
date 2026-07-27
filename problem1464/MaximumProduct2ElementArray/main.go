package main

import "fmt"

//two difference i and j
func maxProduct(nums []int) int {
	num1,num2 := 0,0
	for _,x := range nums{
		if x >= num1{
			num2 = num1
			num1 = x
		}else if x >= num2{
			num2 = x
		}
	}
	return (num1-1)*(num2-1)
}

func main(){
	ans := maxProduct([]int{3,7})
	fmt.Printf("ans: %v\n",ans)
}