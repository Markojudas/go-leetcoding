package main

import "fmt"

func main() {

	//example 1
	//expected output: 15
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))

	//example 2
	//expected output: 6
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}
