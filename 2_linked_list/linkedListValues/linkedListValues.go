package main

type Node struct {
	Val  interface{}
	Next *Node
}

func main () {}

// Iterative Approach - O(n) Runtime
func linkedListValuesIterative(head *Node) []interface{} {
	values := []interface{}{}
	current := head
	for current != nil {
		values = append(values, current.Val)
		current = current.Next
	}
	return values
}

// Recursive Approach - O(n) Runtime
func linkedListValuesRecursive(head *Node) []interface{} {
	if head == nil {
		return []interface{}{}
	}
	values := []interface{}{head.Val}
	return append(values, linkedListValuesRecursive(head.Next)...)
}

