package main

import "fmt"

func countOdds(low int, high int) int {
	var count int

	i := low

	for i <= high {
		if i%2 != 0 {
			count++
		}
		i++
	}

	return count
}

func main() {
	fmt.Println(countOdds(3, 7))
	fmt.Println(countOdds(8, 10))
}
