package main

import (
	"fmt"
)

// func dailyTemperatures(temperatures []int) []int {
// 	stack := []int{}
// 	ans := []int{}
// 	for i,x := range temperatures{
// 		j := i+1
// 		stack = append(stack, temperatures[i])
// 		for {
// 			if j > len(temperatures)-1{
// 				ans = append(ans, 0)
// 				stack = stack[:0]
// 				break
// 			}else if temperatures[j] <= x {
// 				stack = append(stack, temperatures[j])
// 			}else{
// 				ans = append(ans, len(stack))
// 				stack = stack[:0]
// 				break
// 			}
// 			j+=1
// 		}
// 	}
// 	return ans
// }
func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
    ans := make([]int, n)
    stack := []int{} // lưu index

    for i := 0; i < n; i++ {

        for len(stack) > 0 &&
            temperatures[i] > temperatures[stack[len(stack)-1]] {

            idx := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            ans[idx] = i - idx
        }

        stack = append(stack, i)
    }

    return ans
}



func main(){
	temp := []int{30,60,90}
	ans := dailyTemperatures(temp)
	fmt.Printf("Arr : %v\n",ans)
}