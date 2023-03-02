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
	
	a.Next = b
	b.Next = c
	c.Next = d

	x := &Node{Val: "X"}
	y := &Node{Val: "Y"}
	z := &Node{Val: "Z"}

	x.Next = y
	y.Next = z

	zipperListRecursive(a, x)
	fmt.Println(a.Next.Val)
	fmt.Println(a.Next.Next.Val)
	fmt.Println(a.Next.Next.Next.Val)
	fmt.Println(a.Next.Next.Next.Next.Val)
	fmt.Println(a.Next.Next.Next.Next.Next.Val)
	fmt.Println(a.Next.Next.Next.Next.Next.Next.Val)
}

// Iterative Solution
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

// Recursive Solution
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
	fmt.Println(head1.Next)
	fmt.Println(head2.Next)
	next1 := head1.Next
	next2 := head2.Next
	head1.Next = head2
	head2.Next = zipperListRecursive(next1, next2)
	return head1
}