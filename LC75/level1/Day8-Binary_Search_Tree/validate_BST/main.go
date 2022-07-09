package main

import "fmt"

func main() {

	//example 1
	root := createExampleTree1()
	fmt.Println(isValidBST(root)) //true

	//example 2
	root = createExampleTree2()
	fmt.Println(isValidBST(root)) //false

}
