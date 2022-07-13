package main

import (
	"math"
)

func minCostClimbingStairs(cost []int) int {
	//Create an array for the min cost at each step (1 or 2 steps)
	//The array's length should be 1 longer than the length of the cost
	//Because we can treat the "top floor" as a step to reach.

	minCost := make([]int, len(cost)+1)

	//start iteration from step 2, since the mincost of reaching step 0 and 1 is 0.

	for i := 2; i < len(minCost); i++ {
		oneStep := float64(minCost[i-1] + cost[i-1])
		twoStep := float64(minCost[i-2] + cost[i-2])

		minCost[i] = int(math.Min(oneStep, twoStep))
	}

	//the final element in minCost refers to the top floor
	return minCost[len(minCost)-1]
}
