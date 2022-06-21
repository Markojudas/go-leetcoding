package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	res := []int{}

	for _, val := range nums1 {
		found := false
		for i, val2 := range nums2 {
			if val2 == val {
				found = true
			}

			if val2 > val && found {
				res = append(res, val2)
				break
			}

			if i == len(nums2)-1 {
				res = append(res, -1)
			}
		}
	}

	return res
}

func main() {

	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
}
