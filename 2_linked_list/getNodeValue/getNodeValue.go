package main

type Node struct {
	Val  interface{}
	Next *Node
}

func main () {}

// Iterative Approach - O(n) Runtime
func getNodeValueIterative(head *Node, index int) interface{} {
	count := 0
	current := head
	for current != nil {
		if count == index {
			return current.Val
		}
		count ++
		current = current.Next
	}
	return nil
}

// Recursive Approach - O(n) Runtime
func getNodeValueRecursive(head *Node, index int) interface{} {
	if head == nil {
		return nil
	}
	if index == 0 {
		return head.Val
	}
	return getNodeValueRecursive(head.Next, index-1)
}
