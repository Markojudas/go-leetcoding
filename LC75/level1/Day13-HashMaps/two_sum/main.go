package main

import "fmt"

func main() {

	//example 1
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))

	//example 2
	fmt.Println(twoSum([]int{3, 2, 4}, 6))

	//example 3
	fmt.Println(twoSum([]int{3, 3}, 6))

	//fail test-case
	fmt.Println(twoSum([]int{-1, -2, -3, -4, -5}, -8))

}
