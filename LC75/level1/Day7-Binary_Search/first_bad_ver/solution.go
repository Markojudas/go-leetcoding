package main

type product struct {
	states []bool
}

var prod product

// returning true if false; false otherwise
func isBadVersion(version int) bool {

	return !prod.states[version]

}

//helper function to set the test case
func constructor(firstBadVer, n int) {

	prod.states = make([]bool, n+1)

	for ver := range prod.states {
		if ver < firstBadVer {
			prod.states[ver] = true //setting the states; true = good, false(default) = bad
		} else if ver == firstBadVer {
			break
		}
	}
}

func firstBadVersion(n int) int {
	left := 1 //starts with a good version
	right := n

	for left < right {
		mid := left + (right-left)/2 //since we are starting at 1

		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
