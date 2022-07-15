package main

import "fmt"

func main() {

	//example 1
	fmt.Println(backspaceCompare("ab#c", "ad#c"))

	//example 2
	fmt.Println(backspaceCompare("ab##", "c#d#"))

	//example 3
	fmt.Println(backspaceCompare("a#c", "b"))
}
