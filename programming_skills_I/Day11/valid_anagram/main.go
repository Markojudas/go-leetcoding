package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {

	// first check if the strings are of same length
	if len(s) != len(t) {
		return false
	}

	// convert the strings to slice of strings
	sx := strings.Split(s, "")
	tx := strings.Split(t, "")

	// sort the slices
	sort.Strings(sx)
	sort.Strings(tx)

	//compare the slices
	for idx := range sx {
		if sx[idx] != tx[idx] {
			return false
		}
	}

	return true
}

func isAnagramMap(s string, t string) bool {
	// first check if the strings are of same length
	if len(s) != len(t) {
		return false
	}

	//create map, key is a byte, value an int
	dict := make(map[rune]int)

	//populate the map, incrementing the value at index of each byte/char for S, and decrementing for T
	for _, val := range s {
		dict[val]++
	}

	for _, val := range t {
		dict[val]--
	}

	//all the values of the map have to be 0 for it to be an anagram, if it isn't that means there was a different letter

	for key := range dict {

		if dict[key] != 0 {
			return false
		}
	}

	// if the code gets here that means it is an anagram
	return true
}

func main() {

	//example 1:
	fmt.Println(isAnagram("anagram", "nagaram"))

	//example 2:
	fmt.Println(isAnagram("rat", "car"))

	//example 1.2:
	fmt.Println(isAnagramMap("anagram", "nagaram"))

	//example 2.2:
	fmt.Println(isAnagramMap("rat", "car"))
}
