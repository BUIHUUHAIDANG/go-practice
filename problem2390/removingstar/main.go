package main

import "fmt"

type stack []byte

func (s *stack) push(c byte){
	*s = append(*s, c)
}

func (s *stack) pop()  { 
	*s = (*s)[:len(*s)-1]
}

func removeStars(s string) string { 
	var st stack
	word := []byte(s)
	for _,v := range word{
		if v == '*'{
			st.pop()
		}else{
			st.push(v)
		}
	}
	return string(st)
}

func main(){
	fmt.Print(removeStars("erase*****"))
}
