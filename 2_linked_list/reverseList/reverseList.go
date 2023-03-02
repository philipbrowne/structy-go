package main

import "fmt"

type Node struct {
	Val  interface{}
	Next *Node
}

func main () {
	a := &Node{Val: "A"}
	b := &Node{Val: "B"}
	c := &Node{Val: "C"}
	d := &Node{Val: "D"}
	e := &Node{Val: "E"}
	f := &Node{Val: "F"}
	g := &Node{Val: "G"}
	h := &Node{Val: "H"}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	f.Next = g
	g.Next = h

	reverseListRecursively(a)
	fmt.Println(a, b.Next, c.Next, d.Next)
}

// Iterative Solution
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

// Recursive Solution
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