package main

type Node struct {
	Val interface{}
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
}

func removeNodeIterative(head *Node, targetValue interface{}) *Node {
	if head.Val == targetValue {
		return head.Next
	}
	currentNode := head
	var previousNode *Node
	for currentNode != nil {
		if currentNode.Val == targetValue {
			previousNode.Next = currentNode.Next
			break
		}
		previousNode = currentNode
		currentNode = currentNode.Next
	}
	return head
}

func removeNodeRecursive(head *Node, targetValue interface{}) *Node {
	if head == nil {
		return nil
	}
	if head.Val == targetValue {
		return head.Next
	}
	head.Next = removeNodeRecursive(head.Next, targetValue)
	return head
}