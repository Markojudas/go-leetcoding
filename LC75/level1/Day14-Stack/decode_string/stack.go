package main

func newStack() []byte {

	return []byte{}
}

func push(stack []byte, char byte) []byte {

	stack = append(stack, char)

	return stack
}

func peek(stack []byte) rune {

	return rune(stack[len(stack)-1])
}

func pop(stack []byte) ([]byte, byte) {

	char := stack[len(stack)-1]

	stack = stack[:len(stack)-1]

	return stack, char
}

func isEmpty(stack []byte) bool {

	return len(stack) == 0
}
