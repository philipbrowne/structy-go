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
	fmt.Printf("Get node value iterative for index %d: %v\n", 2, getNodeValueIterative(&a, 2))
	fmt.Printf("Get node value recursive for index %d: %v\n", 3, getNodeValueRecursive(&a, 3))
}

// Iterative Solution
func getNodeValueIterative(head *Node, index int) interface{} {
	count := 0
	current := head
	for current != nil {
		if count == index {
			return current.Val
		}
		count ++
		current = current.Next
	}
	return nil
}

// Recursive Solution
func getNodeValueRecursive(head *Node, index int) interface{} {
	if head == nil {
		return nil
	}
	if index == 0 {
		return head.Val
	}
	return getNodeValueRecursive(head.Next, index-1)
}
