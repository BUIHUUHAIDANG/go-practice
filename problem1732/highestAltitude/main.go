package main

import "fmt"

func largestAltitude(gain []int) int {
	sumValue := 0
	maxValue := 0
	for _,v := range gain{
		sumValue+=v
		if sumValue >= maxValue{
			maxValue = sumValue
		}
	}
	return maxValue
}

func main(){
	fmt.Print(largestAltitude([]int{-4,-3,-2,-1,4,3,2}))

}