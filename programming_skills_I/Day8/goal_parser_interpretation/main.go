package main

import "fmt"

func interpret(command string) string {
	res := []byte{}

	xs := []byte(command)

	for i := 0; i < len(xs); i++ {
		if string(xs[i]) == "G" {
			res = append(res, 'G')
		} else if string(xs[i]) == "(" {
			if string(xs[i+1]) != ")" {
				res = append(res, 'a')
				res = append(res, 'l')
				i += 3
			} else {
				res = append(res, 'o')
				i++
			}
		}
	}

	return string(res)
}

func main() {

	//example 1
	fmt.Println(interpret("G()(al)"))

	//example 2
	fmt.Println(interpret("G()()()()(al)"))

	//example 3
	fmt.Println(interpret("(al)G(al)()()G"))

}
