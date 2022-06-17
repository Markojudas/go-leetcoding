package main

import (
	"fmt"
	"math"
)

func average(salary []int) float64 {
	var ans float64

	totalNum := float64(len(salary) - 2)

	smallest := math.MaxFloat64
	highest := math.SmallestNonzeroFloat64

	for _, val := range salary {
		if float64(val) < smallest {
			smallest = float64(val)
		}
		if float64(val) > highest {
			highest = float64(val)
		}

		ans += float64(val)
	}

	return (ans - smallest - highest) / totalNum
}

func main() {

	fmt.Println(average([]int{4000, 3000, 1000, 2000}))
	fmt.Println(average([]int{1000, 2000, 3000}))
}
