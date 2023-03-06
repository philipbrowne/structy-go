package main

import (
	"fmt"
	"math"
)

type Node struct {
	Val float64
	Left *Node
	Right *Node
}

func main () {}

// Depth First Iterative Approach - O(n) Runtime
func treeMinValueDFI(root *Node) float64 {
	minimum := math.Inf(1)
	stack := []*Node{root}
	for len(stack) > 0 {
		node := stack[len(stack) -1]
		stack = stack[:len(stack)-1]
		fmt.Println(node.Val)
		if node.Val < minimum {
			minimum = node.Val
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return minimum
}

// Depth First Recursive Approach - O(n) Runtime
func treeMinValueDFR(root *Node) float64 {
	if root == nil {
		return math.Inf(1)
	}
	minLeft := treeMinValueDFR(root.Left)
	minRight := treeMinValueDFR(root.Right)
	minBoth := math.Min(minLeft, minRight)
	return math.Min(root.Val, minBoth)
}

// Breadth First Approach - O(n) Runtime
func treeMinValueBF(root *Node) float64 {
	minimum := math.Inf(1)
	queue := []*Node{root}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.Val < minimum {
			minimum = current.Val
		}
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
	return minimum
}

