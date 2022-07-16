package main

import (
	"strconv"
	"unicode"
)

func push(stack []byte, char byte) []byte {

	stack = append(stack, char)

	return stack
}

func peek(stack []byte) byte {

	return stack[len(stack)-1]
}

func pop(stack []byte) ([]byte, byte) {

	char := stack[len(stack)-1]

	stack = stack[:len(stack)-1]

	return stack, char
}

func isEmpty(stack []byte) bool {

	return len(stack) == 0
}

func decodeString(s string) string {
	// create stack
	stack := []byte{}

	// range over the string
	for idx := range s {

		//we range until the innermost "]", we are working from the inside out
		if s[idx] == ']' {
			// byte slice to keep the decodedString so far
			decodedString := []byte{}

			//we keep appending to the decodedString until we reach "["
			for peek(stack) != '[' {
				newStack, poppedChar := pop(stack)
				stack = newStack
				decodedString = append(decodedString, poppedChar)
			}

			//we have reached "[", and discard this character
			newStack, _ := pop(stack)
			stack = newStack
			base := 1
			k := 0

			// let's get the value of K
			for !isEmpty(stack) && unicode.IsDigit(rune(peek(stack))) {
				newStack, poppedChar := pop(stack)
				stack = newStack
				intChar, _ := strconv.Atoi(string(poppedChar))
				k = k + intChar*base
				base *= 10
			}

			//pushing the decoded string (reversed) to the stack
			for k != 0 {
				for i := len(decodedString) - 1; i >= 0; i-- {
					stack = push(stack, decodedString[i])
				}
				k--
			}
			// we are building the stack until we we reach the innermost "]"
		} else {
			stack = push(stack, s[idx])
		}
	}

	//let's build the result string by popping the stack
	result := make([]byte, len(stack))
	for idx := len(result) - 1; idx >= 0; idx-- {
		newStack, poppedChar := pop(stack)
		stack = newStack
		result[idx] = poppedChar
	}

	//return the result as a string
	return string(result)
}
