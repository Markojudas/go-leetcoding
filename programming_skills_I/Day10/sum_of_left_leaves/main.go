package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func processSubStree(subtree *TreeNode, isLeft bool) int {
	//preorder traversal of the tree

	//base case
	if subtree.Left == nil && subtree.Right == nil {
		if isLeft {
			return subtree.Val
		} else {
			return 0
		}
	}

	//recursion
	total := 0
	if subtree.Left != nil {
		total += processSubStree(subtree.Left, true)
	}
	if subtree.Right != nil {
		total += processSubStree(subtree.Right, false)
	}

	return total

}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return processSubStree(root, false)
}

func main() {

	// Example 1:
	root1 := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(sumOfLeftLeaves(&root1)) //24

	// Example 2
	root2 := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(sumOfLeftLeaves(&root2)) //0
}
