package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	decimalRep := head.Val

	for head.Next != nil {
		decimalRep = decimalRep*2 + head.Next.Val
		head = head.Next
	}

	return decimalRep
}

func main() {

	//Example 1
	binaryRep := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	fmt.Println(getDecimalValue(&binaryRep))

	//Example 2
	binaryRep2 := ListNode{
		Val:  0,
		Next: nil,
	}
	fmt.Println(getDecimalValue(&binaryRep2))

}
