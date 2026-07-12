package main

import (
	"fmt"
	"slices"
)



func arrayRankTransform(arr []int) []int {
	if len(arr) == 0{
		return []int{}
	}
	tmp := make([]int,len(arr))
	copy(tmp,arr)
	slices.Sort(arr)
	mapRank := make(map[int]int)
	mapRank[arr[0]] =1
	for i := 1 ; i< len(arr); i++{
		if arr[i] == arr[i-1]{
			mapRank[arr[i]] = mapRank[arr[i-1]]
			continue
		}
		mapRank[arr[i]] = mapRank[arr[i-1]] + 1 
	}
	for i := 0 ;i<len(tmp);i++{
		tmp[i] = mapRank[tmp[i]]
	}
	return tmp
}

func main(){
	arr := []int{}
    ans := arrayRankTransform(arr)
	fmt.Printf("Arr: %v",ans)
}