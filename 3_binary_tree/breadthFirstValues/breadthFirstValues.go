package main

type Node struct {
	Val interface{}
	Left *Node
	Right *Node
}

func main () {}

// O(n) Runtime
func breadthFirstValues(root *Node) []interface{} {
	if root == nil {
		return []interface{}{}
	}
	values := []interface{}{}
	queue := []*Node{root}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		values = append(values, current.Val)
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
	return values
}