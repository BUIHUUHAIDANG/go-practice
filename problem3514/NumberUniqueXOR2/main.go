package main

import (
	"fmt"
)


func uniqueXorTriplets(nums []int) int {
	n := len(nums)
	m := 0
	for _, v := range nums {
		m = max(m, v)
	}
	u := 1
	for u <= m {
		u <<= 1
	}
	s := make([]bool, u)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			s[nums[i]^nums[j]] = true
		}
	}
	t := make([]bool, u)
	for x := 0; x < u; x++ {
		if !s[x] {
			continue
		}
		for _, v := range nums {
			t[x^v] = true
		}
	}
	ans := 0
	for _, b := range t {
		if b {
			ans++
		}
	}
	return ans
}
// if len(nums) <=2 => return n
func main(){
	ans := uniqueXorTriplets([]int{4,5,13,17})
	fmt.Print(ans)
}