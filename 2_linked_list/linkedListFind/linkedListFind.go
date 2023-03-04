package main

type Node struct {
	Val  interface{}
	Next *Node
}

func main () {}

// Iterative Approach - O(n) Runtime
func linkedListFindIterative(head *Node, val interface{}) bool {
	current := head
	for current != nil {
		if current.Val == val {
			return true
		}
		current = current.Next
	}
	return false
}

// Recursive Approach - O(n) Runtime
func linkedListFindRecursive(head *Node, val interface{}) bool {
	if head == nil {
		return false
	}
	if head.Val == val {
		return true
	}
	return linkedListFindRecursive(head.Next, val)
}

