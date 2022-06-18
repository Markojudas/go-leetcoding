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

func main() {

	fmt.Println(substractProductAndSum(234))
	fmt.Println(substractProductAndSum(4421))
}
