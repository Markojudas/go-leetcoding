package main

// returning true if false; false otherwise
func isBadVersion(version int) bool {

	return !prod.states[version]

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
