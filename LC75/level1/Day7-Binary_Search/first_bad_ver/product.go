package main

type product struct {
	states []bool
}

var prod product

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
