package main

import "fmt"

func Recursive(p1, p2, i, j int, nums []int, turn bool) bool {
    if i == j {
        if turn {
            p1 += nums[i]
        } else {
            p2 += nums[i]
        }
        return p1 >= p2
    }

    if turn {
        return Recursive(p1+nums[i], p2, i+1, j, nums, false) ||
               Recursive(p1+nums[j], p2, i, j-1, nums, false)
    }

    return Recursive(p1, p2+nums[i], i+1, j, nums, true) &&
           Recursive(p1, p2+nums[j], i, j-1, nums, true)
}

func predictTheWinner(nums []int) bool {
    return Recursive(0, 0, 0, len(nums)-1, nums, true)
}

func main(){
	ans := predictTheWinner([]int{1,5,2})
	fmt.Printf("ans: %v\n",ans)
}