package main

import (
	"fmt"
	"sort"
)

func sortByBits(arr []int) []int {

	sort.Slice(arr, func(i, j int) bool {
		x, y := oneBits(arr[i]), oneBits(arr[j])

		if x == y {
			return arr[i] < arr[j]
		}
		return x < y
	})

	return arr
}

func oneBits(num int) int {
	res := 0

	for num > 0 {
		res += num & 1
		num >>= 1 // if more than 1 digit if will go from left to right.
	}

	return res
}

func main() {

	//example 1
	fmt.Println(sortByBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))

	//example 2
	fmt.Println(sortByBits([]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}))
}
