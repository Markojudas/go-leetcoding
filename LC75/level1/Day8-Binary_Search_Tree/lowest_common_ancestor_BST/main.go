package main

import "fmt"

func main() {

	//example 1
	//expected output: 6
	root, p, q := createTreeExample1()
	fmt.Println(lowestCommonAncestor(root, p, q).Val)

	//example 2
	//expected output: 2
	root, p, q = createTreeExample2()
	fmt.Println(lowestCommonAncestor(root, p, q).Val)

	//example 3
	//expected output: 2
	root, p, q = createTreeExample3()
	fmt.Println(lowestCommonAncestor(root, p, q).Val)
}
