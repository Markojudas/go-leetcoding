package main

import "fmt"

func runningSum(nums []int) []int {

	currentTotal := 0

	for idx, val := range nums {
		nums[idx] = val + currentTotal
		currentTotal += val
	}

	return nums
}

func main() {

	//example 1
	fmt.Println(runningSum([]int{1, 2, 3, 4}))

	//example 2
	fmt.Println(runningSum([]int{1, 1, 1, 1, 1}))

	//example 3
	fmt.Println(runningSum([]int{3, 1, 2, 10, 1}))
}
