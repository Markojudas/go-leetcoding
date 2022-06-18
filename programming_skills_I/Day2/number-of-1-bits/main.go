package main

import "fmt"

// solution
func hammingWeight(num uint32) int {

	// variable to "count" the number of 1's
	var count int

	// variable to go over the 32 bits (unsigned)
	i := 0

	// variable to perform the AND operation (bit manipulation)
	var mask uint32 = 1

	for i < 32 {
		if num&mask != 0 {
			count++
		}

		mask <<= 1 //shifting one position to the left
		i++
	}

	return count
}

func main() {

	fmt.Println(hammingWeight(00000000000000000000000000001011))
	fmt.Println(hammingWeight(00000000000000000000000010000000))

}
