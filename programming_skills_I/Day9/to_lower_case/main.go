package main

import "fmt"

func toLowerCase(s string) string {

	var res []byte

	for idx, val := range []byte(s) {
		if val >= 'A' && val <= 'Z' {
			res = append(res, s[idx]+32)
		} else {
			res = append(res, s[idx])
		}
	}

	return string(res)
}

func main() {

	// Exmaple 1
	fmt.Println(toLowerCase("Hello"))

	// Example 2
	fmt.Println(toLowerCase("here"))

	// Example 3
	fmt.Println(toLowerCase("LOVELY"))

}
