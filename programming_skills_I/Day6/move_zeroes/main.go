package main

import "fmt"

func moveZeroes(nums []int) {
	//first create variables to help
	i := 0

	//identify the zeroes while placing the non-zero values in in their place
	for _, val := range nums {
		if val != 0 {
			nums[i] = val
			i++
		}
	}

	//add the zeros back to the end of the array
	for j := i; j < len(nums); j++ {
		nums[j] = 0
	}

}

func main() {

	nums := []int{0, 1, 0, 3, 12}
	nums2 := []int{0}

	moveZeroes(nums)
	moveZeroes(nums2)

	fmt.Println(nums)
	fmt.Println(nums2)

}
