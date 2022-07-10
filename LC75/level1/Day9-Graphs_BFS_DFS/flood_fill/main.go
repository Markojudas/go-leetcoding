package main

import "fmt"

func main() {

	//example 1
	//expected output: [[2,2,2],[2,2,0],[2,0,1]]
	image, sr, sc, color := createExampleImage1()
	fmt.Print("Input: ")
	printImageGrid(image)
	fmt.Print("Output: ")
	printImageGrid(floodFill(image, sr, sc, color))

	fmt.Println()
	//example 2
	//expected output: [[0,0,0],[0,0,0]]
	image, sr, sc, color = createExampleImage2()
	fmt.Print("Input: ")
	printImageGrid(image)
	fmt.Print("Output: ")
	printImageGrid(floodFill(image, sr, sc, color))
}
