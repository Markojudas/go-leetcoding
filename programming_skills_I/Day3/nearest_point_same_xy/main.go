package main

import (
	"fmt"
	"math"
)

func nearestValidPoint(x int, y int, points [][]int) int {

	distance := math.MaxInt64
	index := -1

	for idx := range points {
		if x == points[idx][0] || y == points[idx][1] {
			manDistance := int(math.Abs(float64(x-points[idx][0])) + math.Abs(float64(y-points[idx][1])))
			if distance > manDistance {
				distance = manDistance
				index = idx
			}
		}
	}
	return index
}

func main() {

	fmt.Println(nearestValidPoint(3, 4, [][]int{
		{1, 2},
		{3, 1},
		{2, 4},
		{2, 3},
		{4, 4},
	})) // 2

	fmt.Println(nearestValidPoint(3, 4, [][]int{{3, 4}})) //0
	fmt.Println(nearestValidPoint(3, 4, [][]int{{2, 3}})) //-1

}
