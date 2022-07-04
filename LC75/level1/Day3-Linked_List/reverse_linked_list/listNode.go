package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printLinkedList(list *ListNode) {
	fmt.Print("[")
	for list != nil {
		if list.Next == nil {
			fmt.Print(list.Val)
		} else {
			fmt.Print(list.Val, ", ")
		}
		list = list.Next
	}
	fmt.Println("]")
}

func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	prev := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return prev

}
