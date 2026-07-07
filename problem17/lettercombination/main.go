package main

import "fmt"



func letterCombinations(digits string) []string {
	var result []string
	phone := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var dfs func(idx int,path string)
	dfs = func(idx int, path string){
		if len(path) == len(digits){
			result = append(result,path)
			return
		}
		char := phone[digits[idx]]
		for _,v := range char {
			dfs(idx+1,path+string(v))
		}
	} 
	dfs(0,"")
	return result

}

func main(){
	ans := letterCombinations("23")
	fmt.Print(ans)
}
