package main

import "fmt"

func printImageGrid(image [][]int) {
	fmt.Print("[")
	for row := range image {
		fmt.Print("[")
		for idx, val := range image[row] {
			if idx != len(image[row])-1 {
				fmt.Print(val, ", ")
			} else {
				fmt.Print(val)
			}
		}
		if row != len(image)-1 {
			fmt.Print("], ")
		} else {
			fmt.Print("]")
		}
	}
	fmt.Println("]")
}

func createExampleImage1() ([][]int, int, int, int) {

	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}

	return image, 1, 1, 2
}

func createExampleImage2() ([][]int, int, int, int) {
	image := [][]int{
		{0, 0, 0},
		{0, 0, 0},
	}

	return image, 0, 0, 0
}
