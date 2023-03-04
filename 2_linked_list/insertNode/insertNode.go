package main

type Node struct {
	Val  interface{}
	Next *Node
}

func main () {}

// Iterative Approach - O(n) runtime
func insertNodeIterative(head *Node, value interface{}, index int) *Node {
	newNode := &Node{
		Val: value,
	}
	if index == 0 {
		newNode.Next = head
		return newNode
	}
	count := 0
	currentNode := head
	for currentNode != nil {
		if index-1 == count {
			newNext := currentNode.Next
			currentNode.Next = newNode
			newNode.Next = newNext
			break
		}
		count++
		currentNode = currentNode.Next
	}
	return head
}

// Recursive Approach - O(n) runtime
func insertNodeRecursive(head *Node, value interface{}, index int) *Node {
	if index == 0 {
		newNode := &Node{
			Val: value,
			Next: head,
		}
		return newNode
	}
	head.Next = insertNodeRecursive(head.Next, value, index-1)
	return head
}