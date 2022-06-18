package main

import (
	"fmt"
	"strconv"
)

func substractProductAndSum(n int) int {

	//Converting the int into a slice of bytes
	strSlice := []byte(strconv.Itoa(n))

	sum, prod := 0, 1

	for _, val := range strSlice {
		intVal, _ := strconv.Atoi(string(val)) //ignoring the error and converting the character to int
		sum += intVal
		prod *= intVal
	}

	return prod - sum
}

func solution(n int) int {
	sum, prod := 0, 1

	for n != 0 {
		x := n % 10 // rightmost digit

		sum += x
		prod *= x

		n /= 10 // stripping away the rightmost digit
	}

	return prod - sum
}

func main() {

	fmt.Println(substractProductAndSum(234))
	fmt.Println(substractProductAndSum(4421))
	fmt.Println(solution(234))
	fmt.Println(solution(4421))
}
