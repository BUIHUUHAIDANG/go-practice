package main

import "fmt"

func isVowels(c byte) bool{
	if c =='a'||c =='e'||c =='i'||c =='o'||c =='u'||c =='A'||c =='E'||c =='I'||c =='O'||c =='U'{
		return true
	}
	return false
}

func reverseVowels(s string) string {
    word := []byte(s)
	i := 0
	j := len(s)-1
	for i<j {
		if isVowels(word[i])&&isVowels(word[j]){
			word[i],word[j] = word[j],word[i]
			i++
			j--
		}
		for !isVowels(word[i]){
			i++
		}
		for !isVowels(word[j]){
			j--
		}
	}
	return string(word)
}

func main(){
	fmt.Print(reverseVowels("leetcode"))

}