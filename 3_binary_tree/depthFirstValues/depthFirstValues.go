package main

type Node struct {
	Val interface{}
	Left *Node
	Right *Node
}

func main () {}

func depthFirstValuesIterative(root *Node) []interface{} {
	if root == nil {
		return []interface{}{}
	}
	values := []interface{}{}
	stack := []*Node{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		values = append(values, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return values
}

func depthFirstValuesRecursive(root *Node) []interface{} {
	if root == nil {
		return []interface{}{}
	}
	result := []interface{}{root.Val}
	leftValues := depthFirstValuesRecursive(root.Left)
	rightValues := depthFirstValuesRecursive(root.Right)
	result = append(result, leftValues...)
	result = append(result, rightValues...)
	return result
}