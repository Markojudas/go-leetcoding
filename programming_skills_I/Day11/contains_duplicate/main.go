package main

import (
	"fmt"
	"sort"
)

func containsDuplicate(nums []int) bool {

	// sorting the slice
	sort.Ints(nums)

	// traverse the slice and check if the index + 1 == index
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}

func containsDuplicateMap(nums []int) bool {

	numsMap := make(map[int]int)

	for _, val := range nums {
		numsMap[val]++

		if numsMap[val] > 1 {
			return true
		}
	}

	/* 	for key := range numsMap {
		if numsMap[key] > 1 {
			return true
		}
	} */

	return false
}

func main() {

	//example 1:
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))

	//exaple 2:
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))

	//example 3:
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))

	// using a map
	//example 1:
	fmt.Println(containsDuplicateMap([]int{1, 2, 3, 1}))

	//exaple 2:
	fmt.Println(containsDuplicateMap([]int{1, 2, 3, 4}))

	//example 3:
	fmt.Println(containsDuplicateMap([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
