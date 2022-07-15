package main

func push(stack []byte, char rune) []byte {
	stack = append(stack, byte(char))

	return stack
}

func pop(stack []byte) []byte {

	if len(stack) > 0 {
		stack = stack[:len(stack)-1]
	}

	return stack
}

func intoEmptyTextEditor(str string) string {

	// after all a string is a slice of bytes
	stack := make([]byte, 0, len(str)) //aka stack

	// range 1 character at a time
	for _, char := range str {

		//if the char is NOT #
		if char != '#' {
			//append to the slice of bytes
			stack = push(stack, char) //aka push
			//else, rework the length of temp Str to accommodate the deletion of #
		} else {
			stack = pop(stack) //aka pop
		}
	}

	//return the slice of bytes as a string
	return string(stack)

}

func backspaceCompare(s string, t string) bool {

	//string compare
	return intoEmptyTextEditor(s) == intoEmptyTextEditor(t)

}
