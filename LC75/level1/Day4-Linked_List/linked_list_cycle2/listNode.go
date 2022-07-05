package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycleMap(head *ListNode) *ListNode {

	// Making a map  Key: ListNode, Value: bool representing whether or not the node has been visited
	visited := make(map[*ListNode]bool)

	node := head

	for node != nil {
		//the node has been visited, return it
		if visited[node] {
			return node
		}

		//the node has NOT been visited
		visited[node] = true
		node = node.Next
	}

	return nil
}

// USING THE FLOYD'S Tortoise and Hare (slow and fast algo)

func getIntersect(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		//intersect!
		if slow == fast {
			return slow
		}
	}

	return nil
}

func detectCycleFloyds(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	//if there is a cycle the slow and fast pointers will intersect at some node
	//let's find it, if not, then there is no cycle
	intersect := getIntersect(head)
	if intersect == nil {
		return nil
	}

	//Phase 2; let's find the entrance of the cycle.
	//using to pointers starting at the beginning and at the point they intersect

	ptr1 := head
	ptr2 := intersect

	for ptr1 != ptr2 {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}

	return ptr1
}
