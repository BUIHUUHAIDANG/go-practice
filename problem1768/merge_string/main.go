package main

import "fmt"


func mergeAlternately(word1 string,word2 string) string {
	var res string
	num := 0
	for num < len(word1) || num < len(word2){
		if num < len(word1){
			res+= string(word1[num])
		}
		if num < len(word2){
			res+= string(word2[num])
		}
		num++
	}
	return res
}

func main(){
	fmt.Print(mergeAlternately("abc","pqr"))
}