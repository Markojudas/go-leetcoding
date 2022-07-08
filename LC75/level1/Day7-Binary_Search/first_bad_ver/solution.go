package main

type versions struct {
	state []bool
}

var versionStates versions

func isBadVersion(version int) bool {

	return versionStates.state[version]

}

//helper function to set the test case
func constructor(setBad, n int) {

	versionStates = versions{
		state: make([]bool, n+1), //plus one because it includes n
	}

	idx := 0

	for idx < n+1 {
		if idx >= setBad {
			versionStates.state[idx] = true //from this index onward we have bad versions.
		}
		idx++
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
