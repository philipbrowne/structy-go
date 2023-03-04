package main

type Node struct {
	Val  int
	Next *Node
}

func main () {}

// Iterative Approach - O(n) Runtime
func linkedListSumIterative(head *Node) int {
	total := 0
	current := head
	for current != nil {
		total += current.Val
		current = current.Next
	}
	return total
}

// Recursive Approach - O(n) Runtime
func linkedListSumRecursive(head *Node) int {
	if head == nil {
		return 0
	}
	return head.Val + linkedListSumRecursive(head.Next)
}

