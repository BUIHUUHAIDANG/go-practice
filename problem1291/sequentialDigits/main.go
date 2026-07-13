package main

import (
	"fmt"
	"math"
	"strconv"
)

var sequenceNum = map[int]int{
	2 : 12,
	3 : 123,
	4 : 1234,
	5 : 12345,
	6 : 123456,
	7 : 1234567,
	8 : 12345678,
	9 : 123456789,
}

var eleven = map[int]int{
	2 : 11,
	3 : 111,
	4 : 1111,
	5 : 11111,
	6 : 111111,
	7 : 1111111,
	8 : 11111111,
	9 : 111111111,
}

var listSequence = make(map[int][]int)

func init(){
	for i := 2; i <= 9;i++{
		num := sequenceNum[i]
		for num % 10 != 0 && num< int(math.Pow(10,float64(i))) {
			listSequence[i] = append(listSequence[i], num)
			num += eleven[i]
		}
	}
}



func sequentialDigits(low int, high int) []int {
	n := len(strconv.Itoa(low))
	h := len(strconv.Itoa(high))
	if low > high && low > sequenceNum[9]{
		return []int{}
	}
	ans := []int{}
	for i := n ; i <=h; i++{
		for _,x := range listSequence[i]{
			if x >= low && x <= high{
				ans = append(ans, x)
			}
		}
	}
	return ans	 	
}

func main(){
	ans := sequentialDigits(15753396,106864044)
	fmt.Printf("Arr: %v\n",ans)
}