package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(str string) bool {

	leftPtr := 0
	rightPtr := len(str) - 1

	//loops until the two pointers meet or pass each other
	for leftPtr < rightPtr {
		leftChar := int32(str[leftPtr])
		rightChar := int32(str[rightPtr])

		//making sure the character pointed by the leftPtr is a letter
		for unicode.IsPunct(leftChar) || unicode.IsSpace(leftChar) {
			leftPtr++
			leftChar = int32(str[leftPtr])
		}

		//making sure the letter is lowercase for comparison
		if unicode.IsUpper(leftChar) {
			leftChar = leftChar + 32
		}

		//making sure the character pointer by the rightPtr is a letter
		for unicode.IsPunct(rightChar) || unicode.IsSpace(rightChar) {
			rightPtr--
			rightChar = int32(str[rightPtr])
		}

		//making sure the letter is lowercase for comparison
		if unicode.IsUpper(rightChar) {
			rightChar = rightChar + 32
		}

		//comparing the letter at the left is equal to the letter at the right of the string
		// if not, then it is not a palindrome
		if leftChar != rightChar {
			return false
		}

		//moving the pointers
		leftPtr++
		rightPtr--
	}

	return true
}

func main() {

	fmt.Println(isPalindrome("Race car!!!!!"))
	fmt.Println(isPalindrome("Hello"))
}
