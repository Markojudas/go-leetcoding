package main

import "fmt"

// helper function: getting the next Int from adding the squares of the digits
func getNext(n int) int {
	totalSum := 0

	for n > 0 {
		temp := n % 10 //rightmost digit
		n = n / 10     //stripping away the rightmost digit

		totalSum += temp * temp //getting square of each digits and adding them
	}

	return totalSum
}

// solution function: using the Floyd's Cycle Finiding Algorithm
func isHappy(n int) bool {

	tortoise := n
	hare := getNext(n)

	//either we reach 1 or we have come full circle
	for hare != 1 && tortoise != hare {
		tortoise = getNext(tortoise)
		hare = getNext(getNext(hare))
	}

	return hare == 1
}

func main() {

	fmt.Println(isHappy(19)) //true
	fmt.Println(isHappy(2))  //false
}
