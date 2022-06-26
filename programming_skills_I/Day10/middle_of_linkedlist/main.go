package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// helper function to print the examples; didn't want to retype this :)
func printList(head *ListNode) {

	fmt.Print("[ ")
	for head != nil {
		if head.Next == nil {
			fmt.Println(head.Val, "]")
		} else {
			fmt.Print(head.Val, ", ")
		}
		head = head.Next
	}

}

func main() {

	// Example 1:
	list := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	res := middleNode(&list)
	printList(res)

	// Example 2:
	list2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}

	res2 := middleNode(&list2)
	printList(res2)

}
