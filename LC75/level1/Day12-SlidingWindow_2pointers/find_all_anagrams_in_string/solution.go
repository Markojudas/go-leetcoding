package main

func findAnagrams(s, p string) []int {
	//using 1 map to test out the anagram
	anagram := map[byte]int{}

	//the size of the window
	np := len(p)

	//slice to return
	res := []int{}

	//set the map with each char/byte of p to 1
	for _, r := range []byte(p) {
		anagram[r]++
	}

	//traverse through s
	for idx, r := range []byte(s) {
		//set the char/byte to - 1 (if it was 0, meaning it is a new char/byte, now it is -1; if it was 1, meaning already mapped it will 0)
		anagram[r]--

		//if it is 0, meaning it was a visited char/byte then remove it from the map
		if anagram[r] == 0 {
			delete(anagram, r)
		}

		//we reach the window limit
		if idx-np >= 0 {
			//we add 1 to the mapping from the left of s
			anagram[s[idx-np]]++

			// if it was mapped before it should be 0 and deleted from the map
			if anagram[s[idx-np]] == 0 {
				delete(anagram, s[idx-np])
			}
		}

		//the map should be empty and we have found the starting point of the anagram.
		//add that index to the slice to return
		if len(anagram) == 0 {
			res = append(res, idx-np+1)
		}
	}

	return res
}
