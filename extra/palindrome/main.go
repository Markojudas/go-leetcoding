package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(str string) bool {

	leftPtr := 0
	rightPtr := len(str) - 1

	for leftPtr < rightPtr {
		leftChar := int32(str[leftPtr])
		rightChar := int32(str[rightPtr])

		for unicode.IsPunct(leftChar) || unicode.IsSpace(leftChar) {
			leftPtr++
			leftChar = int32(str[leftPtr])
		}

		if unicode.IsUpper(leftChar) {
			leftChar = leftChar + 32
		}

		for unicode.IsPunct(rightChar) || unicode.IsSpace(rightChar) {
			rightPtr--
			rightChar = int32(str[rightPtr])
		}

		if unicode.IsUpper(rightChar) {
			rightChar = rightChar + 32
		}

		if leftChar != rightChar {
			return false
		}

		leftPtr++
		rightPtr--
	}

	return true
}

func main() {

	fmt.Println(isPalindrome("Race car!!!!!"))
	fmt.Println(isPalindrome("Hello"))
}
