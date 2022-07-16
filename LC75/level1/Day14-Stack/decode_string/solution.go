package main

import (
	"strconv"
	"unicode"
)

func decodeString(s string) string {
	// create stack
	stack := newStack()
	var poppedChar byte

	// range over the string
	for idx := range s {

		//we range until the innermost "]", we are working from the inside out
		if s[idx] == ']' {
			// byte slice to keep the decodedString so far
			decodedString := []byte{}

			//we keep appending to the decodedString until we reach "["
			for peek(stack) != '[' {
				stack, poppedChar = pop(stack)
				decodedString = append(decodedString, poppedChar)
			}

			//we have reached "[", and discard this character
			stack, _ = pop(stack)
			base := 1
			k := 0

			// let's get the value of K
			for !isEmpty(stack) && unicode.IsDigit(peek(stack)) {
				stack, poppedChar = pop(stack)
				intChar, _ := strconv.Atoi(string(poppedChar))
				k = k + intChar*base
				base *= 10
			}

			//pushing the decoded string (reversed) to the stack k times
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
		stack, poppedChar = pop(stack)
		result[idx] = poppedChar
	}

	//return the result as a string
	return string(result)
}
