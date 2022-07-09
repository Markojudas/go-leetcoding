package main

import "fmt"

func main() {

	//example 1
	leaf1 := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	leaf2 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	root := TreeNode{
		Val:   2,
		Left:  &leaf1,
		Right: &leaf2,
	}

	fmt.Println(isValidBST(&root)) //true

	//example 2
	leaf3 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	leaf4 := TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}

	child1 := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	child2 := TreeNode{
		Val:   4,
		Left:  &leaf3,
		Right: &leaf4,
	}

	root2 := TreeNode{
		Val:   5,
		Left:  &child1,
		Right: &child2,
	}
	fmt.Println(isValidBST(&root2)) //false

}
