package main

func createExampleTree1() *TreeNode {
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

	return &root
}

func createExampleTree2() *TreeNode {
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

	return &root2
}
