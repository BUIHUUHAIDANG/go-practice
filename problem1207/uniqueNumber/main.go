package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	res := make(map[int]int)
	set := make(map[int]struct{})
	for _,v := range arr{
		res[v]+= 1
	}
	for _,x := range res {
		_,ok := set[x]
		if ok == true {
			return false
		}
		set[x] = struct{}{}
	}
	return true
}

func main(){
	fmt.Print(uniqueOccurrences([]int{2,2}))
}