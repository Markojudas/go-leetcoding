package main

import "fmt"

func main() {

	//example 1
	//expected output: 1
	grid := createExampleGrid1()

	fmt.Println(numIslands(grid))

	//example 2
	//expected output 3
	grid2 := createExampleGrid2()

	fmt.Println(numIslands(grid2))

}
