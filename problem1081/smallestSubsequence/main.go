package main

import (
	"fmt"
	_"slices"
	_"strings"
)
// var allPath []string =[]string{}

// func dfs(s []byte,n int,start int,used []bool,path []byte){
// 	if len(path)== n{
// 		allPath = append(allPath, string(path))
// 		return
// 	}
// 	for i:=start;i<len(s);i++{
// 		if used[s[i]-97]==true{
// 			continue
// 		}
// 		used[s[i]-97] = true
// 		dfs(s,n,i+1,used,append(path,s[i]))
// 		used[s[i]-97] =false
// 	}
// }

// func smallestSubsequence(s string) string {
// 	allPath = []string{}
// 	used := make([]bool,26)
// 	tmp := string(s[0])
// 	for i := 0 ; i<len(s);i++{
// 		if !strings.Contains(tmp,string(s[i])){
// 			tmp += string(s[i])
// 		}
// 	}
// 	n := len(tmp)
// 	path := []byte{}
// 	dfs([]byte(s),n,0,used,path)
// 	slices.Sort(allPath)
// 	return allPath[0]
// }
func smallestSubsequence(s string) string {
	left := [26]int{}
	for _, ch := range s {
		left[ch-'a']++
	}
	stack := []byte{}
	inStack := [26]bool{}
	for i := range s {
		ch := s[i]
		if !inStack[ch-'a'] {
			for len(stack) > 0 && ch < stack[len(stack)-1] {
				last := stack[len(stack)-1] - 'a'
				if left[last] == 0 {
					break
				}
				stack = stack[:len(stack)-1]
				inStack[last] = false
			}
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}
		left[ch-'a']--
	}
	return string(stack)
}


func main(){
	ans := smallestSubsequence("ecbacba")
	fmt.Printf("ANS: %v\n",ans)
}