package main

import "fmt"

func checkStraightLine(coordinates [][]int) bool {

	x1 := coordinates[0][0]
	x2 := coordinates[1][0]
	y1 := coordinates[0][1]
	y2 := coordinates[1][1]

	//general equation of a line = ax + by = c
	//a = (y2 - y1)
	//b = (x1 - x2)

	a := (y2 - y1)
	b := (x1 - x2)
	c := a*x1 + b*y1

	for i := 2; i < len(coordinates); i++ {
		findC := a*coordinates[i][0] + b*coordinates[i][1]

		if findC != c {
			return false
		}
	}

	return true
}

func main() {

	coordinates1 := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
		{6, 7},
	}

	coordinates2 := [][]int{
		{1, 1},
		{2, 2},
		{3, 4},
		{4, 5},
		{5, 6},
		{7, 7},
	}

	fmt.Println(checkStraightLine(coordinates1))
	fmt.Println(checkStraightLine(coordinates2))
}
