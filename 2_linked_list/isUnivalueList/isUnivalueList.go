package main

type Node struct {
	Val interface{}
	Next *Node
}

func main () {
}

// Iterative Approach - O(n) Runtime
func isUnivalueListIterative(head *Node) bool {
	current := head
	val := head.Val
	for current != nil {
		if current.Val != val {
			return false
		}
		current = current.Next
	}
	return true
}

// Recursive Approach - O(n) Runtime
func isUnivalueListRecursive(head *Node, prev ...interface{}) bool {
	if head == nil { 
		return true
	} else if  len(prev) == 0 || head.Val == prev[0] {
		return isUnivalueListRecursive(head.Next, head.Val)
	} else {
		return false
	}
}