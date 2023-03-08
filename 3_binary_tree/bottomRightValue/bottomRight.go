package main

type Node struct {
	Val interface{}
	Left *Node
	Right *Node
}

func main() {}

// Breadth First Approach - O(n) Runtime
func bottomRightValue(root *Node) interface{} {
	if root == nil {
		return nil
	}
	var current *Node
	queue := []*Node{root}
	for len(queue) > 0 {
		current = queue[0]
		queue = queue[1:]
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
	return current.Val
}