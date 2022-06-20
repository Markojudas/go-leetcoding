package main

import "fmt"

func arraySign(nums []int) int {
	product := 1

	for _, val := range nums {
		product *= val

		if product == 0 {
			return 0
		} else if product < 0 {
			product = -1
		} else {
			product = 1
		}
	}

	return product
}

func main() {
	fmt.Println(arraySign([]int{-1, -2, -3, -4, 3, 2, 1}))
	fmt.Println(arraySign([]int{1, 5, 0, 2, -3}))
	fmt.Println(arraySign([]int{-1, 1, -1, 1, -1}))
}
