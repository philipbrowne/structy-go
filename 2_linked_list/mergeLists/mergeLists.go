package main

type Node struct {
	Val int
	Next *Node
}

func main () {
}

// Iterative Approach - O(min(n,m)) Runtime where n and m are the lengths of each list
func mergeListsIterative(a, b *Node) *Node {
	current1 := a
	current2 := b
	head := &Node{Val: 0}
	tail := head
	for current1 != nil && current2 != nil {
		if current1.Val < current2.Val {
			tail.Next = current1
			current1 = current1.Next
		} else {
			tail.Next = current2
			current2 = current2.Next
		}
		tail = tail.Next
	}
	if current1 != nil {
		tail.Next = current1
	}
	if current2 != nil {
		tail.Next = current2
	}
	return head.Next
}

// Recursive Approach - O(min(n,m)) Runtime where n and m are the lengths of each list
func mergeListsRecursive(a, b *Node) *Node {
	if a == nil && b == nil {
		return nil
	}
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.Val < b.Val {
		next := a.Next
		a.Next = mergeListsRecursive(next, b)
		return a
	} else {
		next := b.Next
		b.Next = mergeListsRecursive(a, next)
		return b
	}
}