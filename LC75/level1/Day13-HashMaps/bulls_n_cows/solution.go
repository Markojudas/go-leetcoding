package main

import "fmt"

func getHint(secret, guess string) string {

	//create map for all the char in the secret and keep as value the frequency of each
	hmap := map[rune]int{}
	for _, val := range secret {
		hmap[val] += 1
	}

	bulls, cows := 0, 0

	for idx, val := range guess {

		// corresponding characters match
		if val == rune(secret[idx]) {
			//update the bulls
			bulls++

			// update the cows
			// if all val characters from the secret
			// were used up
			if hmap[val] <= 0 {
				cows--
			}
			//corresponding char don't match
		} else {
			//update cows
			if hmap[val] > 0 {
				cows++
			}
		}

		//val character was used
		hmap[val] = hmap[val] - 1
	}

	str := fmt.Sprint(bulls, "A", cows, "B")

	return str
}
