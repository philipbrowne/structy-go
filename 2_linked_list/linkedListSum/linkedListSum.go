package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func main () {
	a := Node{Val: 2}
	b := Node{Val: 8}
	c := Node{Val: 3}
	d := Node{Val: -1}
	e := Node{Val: 7}
	a.Next = &b
	b.Next = &c
	c.Next = &d
	d.Next = &e
	fmt.Println("Iterative Total", linkedListSumIterative(&a))
	fmt.Println("Recursive Total", linkedListSumRecursive(&a))
}

// Iterative Solution
func linkedListSumIterative(head *Node) int {
	total := 0
	current := head
	for current != nil {
		total += current.Val
		current = current.Next
	}
	return total
}

// Recursive Solution
func linkedListSumRecursive(head *Node) int {
	if head == nil {
		return 0
	}
	return head.Val + linkedListSumRecursive(head.Next)
}

