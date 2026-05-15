package main

import(
	"fmt"
)
func findMaxAverage(nums []int, k int) float64 {
	var windowSum float64 = 0

	
	for i := 0; i < k; i++ {
		windowSum += float64(nums[i])
	}

	maxSum := windowSum


	for i := k; i < len(nums); i++ {
		windowSum += float64(nums[i]) - float64(nums[i-k])

		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum / float64(k)
}
func main(){
	res:= findMaxAverage([]int{0,4,0,3,2},1)
	fmt.Printf("%x.5f",res)
}