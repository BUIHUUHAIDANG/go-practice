package main

import "fmt"

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}else if n == 1 || n == 2 {
		return 1
	}
	t1 := 0
	t2 := 1
	t3 := 1
	for n != 2 {
		n--
		tmp := t1+t2+t3
		t1 = t2
		t2 = t3
		t3 = tmp 
	}
	return t3	
}

func main(){
	ans := tribonacci(25)
	fmt.Print(ans)
}