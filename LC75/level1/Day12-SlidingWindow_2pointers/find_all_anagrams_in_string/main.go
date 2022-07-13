package main

import "fmt"

func main() {

	//example 1
	fmt.Println(findAnagrams("cbaebabacd", "abc"))

	//example 2
	fmt.Println(findAnagrams("abab", "ab"))

	//fail test case
	fmt.Println(findAnagrams("af", "be"))

	fmt.Println(findAnagrams("baa", "aa"))

}
