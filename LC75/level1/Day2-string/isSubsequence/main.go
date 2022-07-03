package main

import "fmt"

func isSubsequence(s string, t string) bool {

	//Two pointers & the their bounds
	leftBound, rightBound := len(s), len(t)
	leftPtr, rightPtr := 0, 0

	// when we found a match we move both left and right pointer,
	// if not, then we just move the right pointer
	for leftPtr < leftBound && rightPtr < rightBound {
		if s[leftPtr] == t[rightPtr] {
			leftPtr++
		}

		rightPtr++
	}

	//this ensures we have matched the characters in the order of S
	return leftPtr == leftBound
}

func main() {

	//example 1
	fmt.Println(isSubsequence("abc", "ahbgdc"))

	//example 2
	fmt.Println(isSubsequence("axc", "ahbgdc"))

}
