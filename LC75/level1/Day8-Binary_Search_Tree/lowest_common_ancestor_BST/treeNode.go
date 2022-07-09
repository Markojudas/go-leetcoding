package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createTreeExample1() (r, p, q *TreeNode) {
	leaf1 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	leaf2 := TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}

	subChild1 := TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}

	subChild2 := TreeNode{
		Val:   4,
		Left:  &leaf1,
		Right: &leaf2,
	}

	subChild3 := TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}

	subChild4 := TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}

	child1 := TreeNode{
		Val:   2,
		Left:  &subChild1,
		Right: &subChild2,
	}

	child2 := TreeNode{
		Val:   8,
		Left:  &subChild3,
		Right: &subChild4,
	}

	root := TreeNode{
		Val:   6,
		Left:  &child1,
		Right: &child2,
	}

	return &root, &child1, &child2
}

func createTreeExample2() (r, p, q *TreeNode) {
	leaf1 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	leaf2 := TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}

	subChild1 := TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}

	subChild2 := TreeNode{
		Val:   4,
		Left:  &leaf1,
		Right: &leaf2,
	}

	subChild3 := TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}

	subChild4 := TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}

	child1 := TreeNode{
		Val:   2,
		Left:  &subChild1,
		Right: &subChild2,
	}

	child2 := TreeNode{
		Val:   8,
		Left:  &subChild3,
		Right: &subChild4,
	}

	root := TreeNode{
		Val:   6,
		Left:  &child1,
		Right: &child2,
	}

	return &root, &child1, &subChild2
}

func createTreeExample3() (r, p, q *TreeNode) {

	leaf1 := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	root := TreeNode{
		Val:   2,
		Left:  &leaf1,
		Right: nil,
	}

	return &root, &root, &leaf1
}
