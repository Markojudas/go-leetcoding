package main

func main() {

	leftChild1 := TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}

	leaf1 := TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}

	leaf2 := TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}

	rightChild1 := TreeNode{
		Val:   20,
		Left:  &leaf1,
		Right: &leaf2,
	}

	root := TreeNode{
		Val:   3,
		Left:  &leftChild1,
		Right: &rightChild1,
	}

	ans := levelOrder(&root)

	printAns(ans)
}
