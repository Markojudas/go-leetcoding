package main

import "math"

func maxProfit(prices []int) int {
	//largestDifference = maxProfit
	largestDiff := 0

	//variable to store the minprice of the stock
	minPrice := math.MaxInt64

	for _, val := range prices {
		//update the minPrice
		if val < minPrice {
			minPrice = val
		} else {
			//calculate the largest difference/maxprofit
			currentDiff := float64(val) - float64(minPrice)
			largestDiff = int(math.Max(float64(largestDiff), currentDiff))
		}
	}

	//return the largest difference / max profit
	return largestDiff
}
