package main

import "fmt"

func main() {

	//example 1

	//creating nodes
	node2 := ListNode{
		Val:  2,
		Next: nil,
	}

	node3 := ListNode{
		Val:  0,
		Next: nil,
	}

	node4 := ListNode{
		Val:  -4,
		Next: nil,
	}

	//linking them!
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node2

	//creating and attaching the head of the list
	list1 := ListNode{
		Val:  3,
		Next: &node2,
	}

	//using Map
	cycleNode := detectCycleMap(&list1)

	fmt.Println(cycleNode.Val)

	//using Floyd's Tortoise & Hare algo
	cycleNode2 := detectCycleFloyds(&list1)
	fmt.Println(cycleNode2.Val)

	//EXAMPLE 2:
	//creating nodes
	nodeF1 := ListNode{
		Val:  1,
		Next: nil,
	}

	nodeF2 := ListNode{
		Val:  2,
		Next: &nodeF1,
	}

	//linking them
	nodeF1.Next = &nodeF2

	//using Map
	cycleNode3 := detectCycleMap(&nodeF1)

	fmt.Println(cycleNode3.Val)

	//using Floyd's Tortoise & Hare algo
	cycleNode4 := detectCycleFloyds(&nodeF1)
	fmt.Println(cycleNode4.Val)
}
