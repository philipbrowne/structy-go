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
	fmt.Println(linkedListValuesRecursive(a))
	fmt.Println(linkedListValuesIterative(b))
}

// Iterative Method
func linkedListValuesIterative(head *Node) []interface{} {
	values := []interface{}{}
	current := head
	for current != nil {
		values = append(values, current.Val)
		current = current.Next
	}
	return values
}

// Recursive Method
func linkedListValuesRecursive(head *Node) []interface{} {
	if head == nil {
		return []interface{}{}
	}
	values := []interface{}{head.Val}
	return append(values, linkedListValuesRecursive(head.Next)...)
}

