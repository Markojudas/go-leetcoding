package main

import "fmt"

func main() {

	//example 1
	words := createExample1()
	fmt.Println(topKFrequentWord(words, 2))

	//example 2

	words = createExample2()
	fmt.Println(topKFrequentWord(words, 4))
}
