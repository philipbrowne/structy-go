package main

type Node struct {
	Val interface{}
	Left *Node
	Right *Node
}

func main () {}

// Depth First Iterative - O(n) Runtime
func treeIncludesDFI(root *Node, target interface{}) bool {
	if root == nil {
		return false
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Val == target {
			return true
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return false
}

// Depth First Recursive - O(n) Runtime
func treeIncludesDFR(root *Node, target interface{}) bool {
	if root == nil {
		return false
	}
	if root.Val == target {
		return true
	}
	return treeIncludesDFR(root.Left, target) || treeIncludesDFR(root.Right, target)
}

// Breadth First - O(n) Runtime
func treeIncludesBF(root *Node, target interface{}) bool {
	if root == nil {
		return false
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.Val == target{
			return true
		}
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
	return false
}