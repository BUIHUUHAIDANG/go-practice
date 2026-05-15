package main

func findDifference(nums1 []int, nums2 []int) [][]int {
    A, B := make(map[int]bool, 0), make(map[int]bool, 0)
    for _, v := range nums1 {
        A[v] = true
    }
    for _, v := range nums2 {
        _, ok := A[v]
        if !ok {
            B[v] = true
        }else {
            A[v] = false
        }
    }

    R := make([][]int, 2)
    for k, ok := range A {
        if ok {
            R[0] = append(R[0], k)
        }
    }
    for k, _ := range B {
        R[1] = append(R[1], k)
    }
    return R
}