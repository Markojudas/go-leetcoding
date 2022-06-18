package main

import "fmt"

// solution
func hammingWeight(num uint32) int {
	var count int
	i := 0
	var mask uint32 = 1

	for i < 32 {
		if num&mask != 0 {
			count++
		}

		mask <<= 1
		i++
	}

	return count
}

func main() {

	fmt.Println(hammingWeight(00000000000000000000000000001011))
	fmt.Println(hammingWeight(00000000000000000000000010000000))

}
