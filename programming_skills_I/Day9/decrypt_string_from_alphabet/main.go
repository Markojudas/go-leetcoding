package main

import (
	"fmt"
	"strconv"
)

func freqAlphabets(s string) string {
	var list []byte

	for i := len(s) - 1; i >= 0; i-- {
		var sb []byte

		val := 0

		if s[i] == '#' {
			sb = append(sb, s[i-2])
			sb = append(sb, s[i-1])
			val, _ = strconv.Atoi(string(sb))
			i -= 2
		} else {
			sb = append(sb, s[i])
			val, _ = strconv.Atoi(string(sb))
		}

		list = append(list, byte(val+96))
	}

	var res []byte

	for i := len(list) - 1; i >= 0; i-- {
		res = append(res, list[i])
	}

	return string(res)

}

func main() {

	// example 1:
	fmt.Println(freqAlphabets("10#11#12"))

	// example 2:
	fmt.Println(freqAlphabets("1326#"))

}
