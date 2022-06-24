package main

import "fmt"

func findTheDifference(s string, t string) byte {

	var asciiS byte
	var asciiT byte

	for i := 0; i < len(t); i++ {
		if i < len(s) {
			asciiS += s[i]
		}

		asciiT += t[i]
	}

	return asciiT - asciiS
}

func main() {

	fmt.Println(string(findTheDifference("abcd", "abcde")))
}
