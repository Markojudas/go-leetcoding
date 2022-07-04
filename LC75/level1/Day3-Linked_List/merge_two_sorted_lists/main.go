package main

import "fmt"

func main() {

	list1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	list2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	listPtr := mergeTwoLists(&list1, &list2)

	fmt.Print("[")
	for listPtr != nil {

		if listPtr.Next != nil {
			fmt.Print(listPtr.Val, ", ")
		} else {
			fmt.Print(listPtr.Val)
		}

		listPtr = listPtr.Next
	}
	fmt.Println("]")
}
