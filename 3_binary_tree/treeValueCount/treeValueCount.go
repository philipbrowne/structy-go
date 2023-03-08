package main

type Node struct {
	Val interface{}
	Left *Node
	Right *Node
}

func main () {}

// Depth First Recursive Approach - O(n) Runtime
func treeValueCount_DFR(root *Node, target interface{}) int {
	if root == nil {
		return 0
	}
	count := 0
	if root.Val == target {
		count++
	}
	countLeft := treeValueCount_DFR(root.Left, target)
	countRight := treeValueCount_DFR(root.Right, target)
	return count + countLeft + countRight
}

// Depth First Iterative Approach - O(n) Runtime
func treeValueCount_DFI(root *Node, target interface{}) int {
	if root == nil {
		return 0
	}
	count := 0
	stack := []*Node{}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Val == target {
			count++
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return count
}

// Breadth First Approach - O(n) Runtime
func treeValueCount_BF(root *Node, target interface{}) int {
	if root == nil {
		return 0
	}
	count := 0
	queue := []*Node{root}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.Val == target {
			count++
		}
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
	return count
}