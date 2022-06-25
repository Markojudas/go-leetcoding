package main

import "fmt"

func isAlienSorted(words []string, order string) bool {
	orderMap := make([]int, 26)

	for i := 0; i < len(order); i++ {
		orderMap[order[i]-'a'] = i
	}

	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words[i]); j++ {
			if j >= len(words[i+1]) {
				return false
			}

			if words[i][j] != words[i+1][j] {
				currentWordChar := words[i][j] - 'a'
				nextWordChar := words[i+1][j] - 'a'

				if orderMap[currentWordChar] > orderMap[nextWordChar] {
					return false
				} else {
					break
				}
			}

		}
	}

	return true
}

func main() {

	// Example 1:

	words := []string{
		"hello",
		"leetcode",
	}

	order := "hlabcdefgijkmnopqrstuvwxyz"

	fmt.Println(isAlienSorted(words, order))

	// Example 2:

	words2 := []string{
		"word",
		"world",
		"row",
	}

	order2 := "worldabcefghijkmnpqstuvxyz"

	fmt.Println(isAlienSorted(words2, order2))
}
