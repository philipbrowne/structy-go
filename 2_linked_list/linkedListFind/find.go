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
	fmt.Println(`Iterative for "B": `, linkedListFindIterative(&a, "B"))
	fmt.Println(`Recursive for "E": `, linkedListFindRecursive(&a, "E"))
}

// Iterative Solution
func linkedListFindIterative(head *Node, val interface{}) bool {
	current := head
	for current != nil {
		if current.Val == val {
			return true
		}
		current = current.Next
	}
	return false
}

// Recursive Solution
func linkedListFindRecursive(head *Node, val interface{}) bool {
	if head == nil {
		return false
	}
	if head.Val == val {
		return true
	}
	return linkedListFindRecursive(head.Next, val)
}

