package main

type Node struct {
	Val int
	Left *Node
	Right *Node
}

func main () {}

// Iterative Depth First Approach - O(n) Runtime
func treeSumIterativeDF(root *Node) int {
	if root == nil {
		return 0
	}
	sum := 0
	stack := []*Node{root}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		sum += current.Val
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}
	return sum
}

// Recursive Depth-First Approach - O(n) Runtime
func treeSumRecursiveDF(root *Node) int {
	if root == nil {
		return 0
	}
	return root.Val + treeSumRecursiveDF(root.Left) + treeSumRecursiveDF(root.Right)
}

func treeSumBF(root *Node) int {
	if root == nil {
		return 0
	}
	sum := 0
	queue := []*Node{root}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		sum += current.Val
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
	return sum
}