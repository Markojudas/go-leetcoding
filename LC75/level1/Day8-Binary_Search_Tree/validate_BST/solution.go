package main

func isValidBST(root *TreeNode) bool {

	return validate(root, nil, nil)
}

func validate(root *TreeNode, low, high interface{}) bool {

	//empty trees are valid
	if root == nil {
		return true
	}

	// the current node needs to be between low and high
	if (low != nil && root.Val <= low.(int)) || (high != nil && root.Val >= high.(int)) {
		return false
	}

	//validate the left and right subtrees
	return validate(root.Left, low, root.Val) && validate(root.Right, root.Val, high)
}
