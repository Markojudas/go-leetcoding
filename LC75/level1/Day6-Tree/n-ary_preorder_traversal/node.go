package main

type Node struct {
	Val      int
	Children []*Node
}

func getVal(node *Node, results *[]int) {

	if node == nil {
		return
	}

	*results = append(*results, node.Val)

	if node.Children == nil {
		return
	}

	for _, child := range node.Children {
		getVal(child, results)
	}

}

func preorder(root *Node) []int {

	if root == nil {
		return []int{}
	}

	var results = []int{}

	getVal(root, &results)

	return results
}
