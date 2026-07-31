package main

import (
	"fmt"
	"sort"
)

func minimumPushes(word string) int {
	ListCh := make([]int,26)
	ans := 0
	cnt := 0
	for _,x := range word {
		ListCh[int(x-'a')] += 1
	}
	sort.Slice(ListCh,func(i, j int) bool {
		return ListCh[i] > ListCh[j]
	})
	for i := 0; i<26; i++{
		if cnt < 8 && ListCh[i] !=0 {
			ans += 1*ListCh[i]
			cnt +=1
		}else if cnt < 16 && ListCh[i] !=0{
			ans += 2*ListCh[i]
			cnt +=1
		}else if cnt < 24 && ListCh[i] !=0{
			ans += 3*ListCh[i]
			cnt +=1
		}else if cnt <26 && ListCh[i] !=0{
			ans += 4*ListCh[i]
			cnt +=1
		}
	}
	return ans 
}

func main(){
	ans := minimumPushes("aabbccddeeffgghhiiiiii")
	fmt.Printf("ans : %v\n",ans)
}