package main

import "fmt"

func singleNumber(nums []int) int {
    ans := 0
	for _, x := range nums {
    	ans ^= x
	}
	return ans
}

func main(){
	ans := singleNumber([]int{2,3,2,4,4})
	fmt.Print(ans)
}