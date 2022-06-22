package main

import "fmt"

func sumOddLengthSubarrays(arr []int) int {

	// variable to store the sum
	sum := 0

	for i := 0; i < len(arr); i++ {
		idx := 1
		tempSum := 0
		for j := i; j < len(arr); j++ {
			tempSum += arr[j]
			if idx%2 == 1 {
				sum += tempSum
			}
			idx++
		}
	}

	return sum
}

func main() {

	fmt.Println(sumOddLengthSubarrays([]int{1, 4, 2, 5, 3}))
	fmt.Println(sumOddLengthSubarrays([]int{1, 2}))
	fmt.Println(sumOddLengthSubarrays([]int{10, 11, 12}))
}
