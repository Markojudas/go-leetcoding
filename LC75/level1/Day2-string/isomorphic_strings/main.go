package main

import "fmt"

func isIsomorphic(s string, t string) bool {

	//create "dictionaries"
	StoT := make(map[byte]byte)
	TtoS := make(map[byte]byte)

	//setting each index to 0 // NULL to represent empty;
	for idx := range s {
		StoT[s[idx]] = 0
		TtoS[t[idx]] = 0

	}
	for idx := range s {

		c1 := s[idx]
		c2 := t[idx]

		//new mappings
		if StoT[c1] == 0 && TtoS[c2] == 0 {
			StoT[c1] = c2
			TtoS[c2] = c1
		} else if !(StoT[c1] == c2 && TtoS[c2] == c1) {
			//either mapping doesn't exist in one of the dic or mapping exists and it doesn't match in either or both
			return false
		}
	}
	return true
}

func main() {

	//example 1
	fmt.Println(isIsomorphic("egg", "add"))

	//example 2
	fmt.Println(isIsomorphic("foo", "bar"))

	//example 3
	fmt.Println(isIsomorphic("paper", "title"))
}
