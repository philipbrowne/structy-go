package main

type Node struct {
	Val  interface{}
	Next *Node
}

func main () {}

// Iterative Approach - O(n) Runtime
func reverseListIteratively(head *Node) *Node {
	current := head
	var prev *Node
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		if next == nil {
			head = current
		}
		current = next
	}
	return head
}

// Recursive Approach - O(n) Runtime
func reverseListRecursively(head *Node, prevNodes ...*Node) *Node {
	var prev *Node
	if len(prevNodes) > 0 {
		prev = prevNodes[0]
	}
	if head == nil {
		return prev
	}
	next := head.Next
	head.Next = prev
	prev = head
	return reverseListRecursively(next, prev)
}