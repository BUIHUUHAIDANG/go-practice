package main

import "fmt"


func gcd(a int,b int)int{
	for b!=0{
		r := a % b
		a = b
		b = r
	}
	return  a
}

func findGCD(nums []int) int {
	nm := 1001
	nx := 0
	for _,x := range nums{
		nm = min(nm,x)
		nx = max(nx,x)
	}	
	return gcd(nm,nx)
}

func main(){
	ans := findGCD([]int{3,3})
	fmt.Printf("ans: %v\n",ans)
}