package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printAns(ans [][]int) {

	fmt.Print("[")
	for idx1, val1 := range ans {
		fmt.Print("[")
		for idx, val2 := range val1 {
			if idx == len(val1)-1 {
				fmt.Print(val2)
			} else {
				fmt.Print(val2, ", ")
			}
		}
		if idx1 == len(ans)-1 {
			fmt.Print("]")
		} else {
			fmt.Print("],")
		}
	}
	fmt.Println("]")
}

func getVal(node *TreeNode, arr *[][]int, level int) {

	//base condition
	if node == nil {
		return
	}

	//creating a new list at each level
	if len(*arr) == level {
		*arr = append(*arr, []int{})
	}

	//getting the val
	(*arr)[level] = append((*arr)[level], node.Val)

	//recursion
	getVal(node.Left, arr, level+1)
	getVal(node.Right, arr, level+1)

}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	getVal(root, &res, 0)

	return res
}
