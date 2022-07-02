package main

import "fmt"

func pivotIndex(nums []int) int {

	sum, leftsum := 0, 0

	// getting the sum of all the values of the array
	for _, val := range nums {
		sum += val
	}

	//getting the leftsum to find the pivot
	for idx, val := range nums {
		if leftsum == sum-leftsum-val {
			return idx
		}
		leftsum += val
	}

	return -1
}

func main() {
	//example 1
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6})) //index  3

	//example 2
	fmt.Println(pivotIndex([]int{1, 2, 3})) // -1

	//example 3
	fmt.Println(pivotIndex([]int{2, 1, -1})) //index 0; there are no elements at the left of index 0 and to the right (1 -1) = 0

}
