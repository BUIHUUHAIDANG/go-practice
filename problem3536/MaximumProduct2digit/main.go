package main

import "fmt"

func maxProduct(n int) int {
	max1 := 0
	max2 := 0
	for n != 0 {
		r := n % 10
		if r >= max1{
			max2 = max1
			max1 = r
		}
		if max1 > r && r >= max2{
			max2 = r
		}
		n = n / 10
	}
	return max1*max2
}

func main(){
	ans := maxProduct(124)
	fmt.Printf("ANS : %v\n",ans)
}