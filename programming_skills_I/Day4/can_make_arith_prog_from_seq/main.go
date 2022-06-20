package main

import (
	"fmt"
	"sort"
)

func canMakeArithmeticProgressive(arr []int) bool {

	// if the array is of size 2 then it is true
	if len(arr) == 2 {
		return true
	}

	//if not, sort the array
	sort.Ints(arr)

	// calculate the difference
	diff := arr[1] - arr[0]

	// now iterate through the array making sure the difference is the same throughout two consecutive elements

	for i := 1; i < len(arr); i++ {
		if arr[i-1]+diff != arr[i] {
			return false
		}
	}
	return true
}

func main() {

	fmt.Println(canMakeArithmeticProgressive([]int{3, 5, 1})) // true
	fmt.Println(canMakeArithmeticProgressive([]int{1, 2, 4})) // false

}
