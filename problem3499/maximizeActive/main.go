package main

import (
	"fmt"
)


const (
	Active = '1'
	InActive = '0'
)

func maxActiveSectionsAfterTrade(s string) int {
	cntOne :=0
	cntZero := 0
	prv := -1 
	cur := -1 
	mx := 0
	i,j := 0,0
	for i<len(s) || j <len(s){
		for i<len(s) && s[i] == InActive{
			i++
			cntZero += 1
		}
		if cur == -1 && prv == -1 && cntZero != 0{
			cur = cntZero
			cntZero = 0
		}else if cur != -1 && prv == -1 {
			prv = cur
			cur = cntZero
			mx = max(mx,cur+prv)
			prv = -1
			cntZero = 0 
		}
		j = i 
		for j <len(s) && s[j] == Active{
			cntOne += 1
			j++
		}
		i = j 
	}
	if cntOne == 0 {
		return 0
	} 
	return cntOne + mx
}

func main(){
	res := maxActiveSectionsAfterTrade("01101001")
	fmt.Printf("Ans : %v\n",res)
}