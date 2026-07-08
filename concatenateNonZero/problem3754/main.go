package main

import (
	"fmt"
	"math"
)



func sumAndMultiply(n int) int64 {
	if n == 0 {
		return 0
	}
	tmp := n
	modifyNum := 0
	sumNum := 0
	i := 0.0
	for tmp != 0{
		r := tmp % 10
		if r != 0{
			modifyNum = r*int(math.Pow(10,i)) + modifyNum
			sumNum += r
			i++
		}
		tmp = tmp / 10
	}
	
	return int64(sumNum*modifyNum)
}

func main(){
	fmt.Print(sumAndMultiply(10203004))
}