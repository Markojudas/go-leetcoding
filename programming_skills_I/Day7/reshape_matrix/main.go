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
	res := make([][]int, r)

	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}

	if len(mat) == 0 || r*c != len(mat)*len(mat[0]) {
		return mat
	}

	customeQueue := &customQueue{
		queue: []int{},
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			customeQueue.Enqueue(mat[i][j])
		}
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			val, _ := customeQueue.Front()
			customeQueue.Dequeue()
			res[i][j] = val
		}
	}

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
