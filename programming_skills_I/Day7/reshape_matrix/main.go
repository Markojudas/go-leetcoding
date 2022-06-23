package main

import (
	"fmt"
	"math"
)

type customQueue struct {
	queue []int
}

func (c *customQueue) Enqueue(item int) {
	c.queue = append(c.queue, item)
}

func (c *customQueue) Dequeue() error {
	if len(c.queue) > 0 {
		c.queue = c.queue[1:]
		return nil
	}

	return fmt.Errorf("error popping queue")
}

func (c *customQueue) Front() (int, error) {
	if len(c.queue) > 0 {
		return c.queue[0], nil
	}

	return math.MinInt64, fmt.Errorf("peep error")
}

func matrixReshape(mat [][]int, r int, c int) [][]int {

	//	making sure the request is possible and legal
	if len(mat) == 0 || r*c != len(mat)*len(mat[0]) {
		return mat
	}

	// START:
	// creating a 2D int slice with r rows
	res := make([][]int, r)

	// initiating C columns for each row
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}
	// END

	// creating a Queue
	customeQueue := &customQueue{
		queue: []int{},
	}

	// adding each int from the matrix to the queue in order traversal
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			customeQueue.Enqueue(mat[i][j])
		}
	}

	// popping the queue into the 2D int slice
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			val, _ := customeQueue.Front()
			customeQueue.Dequeue()
			res[i][j] = val
		}
	}

	//without using a queue; just creating a slice and appending with he values of the matrix
	/* 	tempSlice := []int{}

	   	for i := 0; i < len(mat); i++ {
	   		for j := 0; j < len(mat[0]); j++ {
	   			tempSlice = append(tempSlice, mat[i][j])
	   		}
	   	}

	   	tempIdx := 0

	   	for i := 0; i < r; i++ {
	   		for j := 0; j < c; j++ {
	   			res[i][j] = tempSlice[tempIdx]
	   			tempIdx++
	   		}
	   	}
	*/
	return res
}

func main() {

	fmt.Println(matrixReshape(
		[][]int{
			{1, 2},
			{3, 4},
		},
		1, 4,
	))

	fmt.Println(matrixReshape(
		[][]int{
			{1, 2},
			{3, 4},
		},
		2, 4,
	))
}
