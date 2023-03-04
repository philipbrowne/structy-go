package main

type Node struct {
	Val interface{}
	Next *Node
}

func main () {
}

// Iterative Approach - O(n) Runtime
func removeNodeIterative(head *Node, targetValue interface{}) *Node {
	if head.Val == targetValue {
		return head.Next
	}
	currentNode := head
	var previousNode *Node
	for currentNode != nil {
		if currentNode.Val == targetValue {
			previousNode.Next = currentNode.Next
			break
		}
		previousNode = currentNode
		currentNode = currentNode.Next
	}
	return head
}

// Recursive Approach - O(n) Runtime
func removeNodeRecursive(head *Node, targetValue interface{}) *Node {
	if head == nil {
		return nil
	}
	if head.Val == targetValue {
		return head.Next
	}
	head.Next = removeNodeRecursive(head.Next, targetValue)
	return head
}