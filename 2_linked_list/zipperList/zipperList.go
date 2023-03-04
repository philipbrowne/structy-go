package main

type Node struct {
	Val  interface{}
	Next *Node
}

func main () {}

// Iterative Approach - O(min(n,m)) Runtime where n and m are the lengths of each list
func zipperListIterative(head1 *Node, head2 *Node) *Node {
	current := head1
	opposite := head2
	for current != nil && opposite != nil {
		curNext := current.Next
		current.Next = opposite
		opposite = curNext
		current = current.Next
	}
	return head1
}

// Recursive Solution - O(min(n,m)) Runtime where n and m are the lengths of each list
func zipperListRecursive (head1 *Node, head2 *Node) *Node {
	if head1 == nil && head2 == nil {
		return nil
	}
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	next1 := head1.Next
	next2 := head2.Next
	head1.Next = head2
	head2.Next = zipperListRecursive(next1, next2)
	return head1
}