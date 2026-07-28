package main

import (
	"fmt"
	"slices"
)


func smallestPalindrome(s string) string {
	if len(s) == 1{
		return s
	}
	mid := len(s)/2
	tmp := []rune(s[:mid])
	slices.Sort(tmp)
	str := string(tmp)
	slices.Reverse(tmp)

	if len(s) % 2 == 0{
		return str+string(tmp)
	}
	return str + string(s[mid])+ string(tmp)
}

func main(){
	ans := smallestPalindrome("rur")
	ap := make([]int,26)
	fmt.Println(len(ap))
	fmt.Printf("ans: %v\n",ans)
}