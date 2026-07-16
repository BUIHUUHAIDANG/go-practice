package main

import "fmt"

func gcd(a int,b int)int{
	for b !=0{
		r := a%b
		a = b
		b = r
	}
	return a
}


func gcdOfOddEvenSums(n int) int {
	return gcd(n*n,n*(n+1))
}

// sum of n num odd is n*n
// sum of n num even is n*(n+1)
// gcd(n,n+1) == 1 
func main(){
	ans := gcdOfOddEvenSums(4)
	fmt.Printf("ans: %v\n",ans)
}