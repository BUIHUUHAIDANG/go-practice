package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxValue := 0
	for i:= range candies {
		if candies[i]>=maxValue{
			maxValue = candies[i]
		}
	}
	res := make([]bool,len(candies))
	for i:= range candies {
		if candies[i]+extraCandies >= maxValue {
			res[i] = true
		}else{
			res[i] = false
		}
	}
	return res
}

func main(){
	candies := []int{2,3,5,1,5}
	extraCandies := 3
	fmt.Print(kidsWithCandies(candies,extraCandies))
}