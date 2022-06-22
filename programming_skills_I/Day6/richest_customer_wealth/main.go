package main

import (
	"fmt"
	"math"
)

func maxiumWealth(accounts [][]int) int {

	// start with the min int64 for comparison
	maxWealth := math.MinInt64

	//ranging on customers
	for _, customer := range accounts {
		custTotalAmount := 0

		//ranging on customer's amount
		for _, amount := range customer {
			custTotalAmount += amount
		}

		// pick the richest!
		if custTotalAmount > maxWealth {
			maxWealth = custTotalAmount
		}
	}

	return maxWealth
}

func main() {

	//example 1
	fmt.Println(maxiumWealth([][]int{
		{1, 2, 3},
		{3, 2, 1},
	}))

	//example 2
	fmt.Println(maxiumWealth([][]int{
		{1, 5},
		{7, 3},
		{3, 5},
	}))

	//example 3
	fmt.Println(maxiumWealth([][]int{
		{2, 8, 7},
		{7, 1, 3},
		{1, 9, 5},
	}))

}
