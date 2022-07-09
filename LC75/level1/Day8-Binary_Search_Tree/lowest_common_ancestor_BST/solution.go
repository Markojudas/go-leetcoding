package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// value p and q
	pVal, qVal := p.Val, q.Val

	//start at root
	node := root

	//traverse the tree:
	for node != nil {

		//value of ancestor/parent node:
		parentVal := node.Val

		if pVal > parentVal && qVal > parentVal {
			//both p and q are greater than parent
			node = node.Right
		} else if pVal < parentVal && qVal < parentVal {
			//both p and q are lesser than parent
			node = node.Left
		} else {
			// we have found the split point, i.e., the  LCA node.
			return node
		}

	}
	return nil
}
