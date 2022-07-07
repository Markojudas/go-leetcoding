package main

import "fmt"

func main() {

	leaf5 := Node{
		Val:      5,
		Children: []*Node{},
	}

	leaf6 := Node{
		Val:      6,
		Children: []*Node{},
	}

	child3 := Node{
		Val: 3,
		Children: []*Node{
			&leaf5,
			&leaf6,
		},
	}

	leaf2 := Node{
		Val:      2,
		Children: []*Node{},
	}

	leaf4 := Node{
		Val:      4,
		Children: []*Node{},
	}

	root := Node{
		Val: 1,
		Children: []*Node{
			&child3,
			&leaf2,
			&leaf4,
		},
	}

	fmt.Println(preorder(&root))
}
