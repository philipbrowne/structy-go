package main

import "fmt"

type Node struct {
	Val  interface{}
	Next *Node
}

func main () {
	a := Node{Val: "A"}
	b := Node{Val: "B"}
	c := Node{Val: "C"}
	d := Node{Val: "D"}
	a.Next = &b
	b.Next = &c
	c.Next = &d
}

// Printing Linked List Iteratively
func printListIteratively(head Node) {
	current := head
	for {
		fmt.Println(current.Val)
		if current.Next == nil {
			break
		}
		current = *current.Next
	}
}

// Printing Linked List Recursively
func printListRecursive(head Node) {
	if &head == nil { 
		return
	}
	fmt.Println(head.Val)
	if head.Next != nil {
		printListRecursive(*head.Next)
	}
}