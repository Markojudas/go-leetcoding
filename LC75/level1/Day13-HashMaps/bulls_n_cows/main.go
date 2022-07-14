package main

import "fmt"

func main() {

	//example 1
	//expected output: "1A3B"
	secret := createExample1Secret()
	fmt.Println(getHint(secret, "7810"))

	//example 2
	//expected output: "1A1B"
	secret = createExample2Secret()
	fmt.Println(getHint(secret, "0111"))
}
