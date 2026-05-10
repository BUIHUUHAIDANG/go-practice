package main

import "fmt"

func gcd(a int,b int) int {
	if b == 0 {
		return a
	}
	return gcd(b,a%b)
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	var n int = len(str1)
	var m int = len(str2)
	var res int = gcd(n,m)
	return str1[:res]
}

func main() {
	fmt.Print(gcdOfStrings("ABCABC","ABC"))

}