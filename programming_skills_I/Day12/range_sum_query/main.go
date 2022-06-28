package main

import "fmt"

type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	na := NumArray{
		arr: nums,
	}

	return na
}

func (nA *NumArray) SumRange(left int, right int) int {
	sum := 0

	for _, val := range nA.arr[left : right+1] {
		sum += val
	}

	return sum
}

func main() {
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})

	fmt.Println(obj.SumRange(0, 2))
	fmt.Println(obj.SumRange(2, 5))
	fmt.Println(obj.SumRange(0, 5))

}
