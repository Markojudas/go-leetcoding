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

func middleNode(head *ListNode) *ListNode {
	// Using Fast & Slow Pointer Algo

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
