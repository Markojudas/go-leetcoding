package main

import "fmt"

func areAlmostEqual(s1 string, s2 string) bool {

	//first check if the 2 strings are equal
	if s1 == s2 {
		return true
	}

	//if not...

	//keep a counter of the num of difference characters
	diff := 0

	//flag for the first and second positions where they are different
	first, last := 0, 0

	//find the existence of the different characters
	//let's create 2 byte slices for comparison
	s1Bs := []byte(s1)
	s2Bs := []byte(s2)

	// since both strings are of same length ranging over either one is fine
	for idx := range s1Bs {

		if s1Bs[idx] != s2Bs[idx] {
			//locating the first different character
			if diff == 0 {
				first = idx
			} else {
				last = idx
			}
			diff++ //updating how many different characters we encounter
		}
	}

	// no difference or if only two characters are different (one swap)
	return (diff == 0 || diff == 2 && (s1Bs[first] == s2Bs[last] && s1Bs[last] == s2Bs[first]))
}

func main() {
	fmt.Println(areAlmostEqual("bank", "kanb"))     //true
	fmt.Println(areAlmostEqual("attack", "defend")) //false
	fmt.Println(areAlmostEqual("kelb", "kelb"))     //true
}
