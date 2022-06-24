package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	var res []byte

	iter1, iter2 := 0, 0

	for iter1 < len(word1) || iter2 < len(word2) {

		if iter1 < len(word1) {
			res = append(res, word1[iter1])
			iter1++
		}

		if iter2 < len(word2) {
			res = append(res, word2[iter2])
			iter2++
		}
	}

	return string(res)
}

func main() {

	//example 1
	fmt.Println(mergeAlternately("abc", "pqr"))

	//exemaple 2
	fmt.Println(mergeAlternately("ab", "pqrs"))

	//example 3
	fmt.Println(mergeAlternately("abcd", "pq"))

}
