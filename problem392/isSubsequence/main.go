package main

import "fmt"

func isSubsequence(s string, t string) bool {
	i:=0
	j:=0
	for i<len(s) && j <len(t){
		if s[i] == t[j]{
			i++
		}
		j++
	}
	if i != len(s){
		return false
	}
	return true
}

func main(){
	fmt.Print(isSubsequence("axc","ahbgdc"))
}