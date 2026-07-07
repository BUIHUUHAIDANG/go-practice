package main

import "fmt"

func partition(nums []int, low, high int) int {
	pivot := nums[low]

	i := low - 1
	j := high + 1

	for {
		for {
			i++
			if nums[i] >= pivot {
				break
			}
		}

		for {
			j--
			if nums[j] <= pivot {
				break
			}
		}

		if i >= j {
			return j
		}

		nums[i], nums[j] = nums[j], nums[i]
	}
}

func findKthLargest(nums []int, k int) int {
	target := len(nums) - k

	left := 0
	right := len(nums) - 1

	for left < right {
		p := partition(nums, left, right)

		if target <= p {
			right = p
		} else {
			left = p + 1
		}
	}

	return nums[target]
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2

	fmt.Println(findKthLargest(nums, k)) // 5
}