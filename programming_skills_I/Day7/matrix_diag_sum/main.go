package main

import "fmt"

func diagonalSum(mat [][]int) int {

	var sum int

	secondaryIdx := len(mat[0]) - 1

	for idx := range mat {

		if idx == secondaryIdx {
			sum += mat[idx][secondaryIdx]
		} else {
			sum += mat[idx][idx] + mat[idx][secondaryIdx]
		}
		secondaryIdx--
	}

	return sum
}

func main() {

	//example 1:
	fmt.Println(diagonalSum([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))

	//example 2:
	fmt.Println(diagonalSum([][]int{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}))

	//example 3:
	fmt.Println(diagonalSum([][]int{
		{5},
	}))
}
