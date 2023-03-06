package main

import "math"

type Node struct {
	Val float64
	Left *Node
	Right *Node
}

func main () {}

// Recursive Depth First Approach - O(n) Runtime
func maxPathSum(root *Node) float64 {
	if root == nil {
		return math.Inf(-1)
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	return math.Max(root.Val + maxPathSum(root.Left), root.Val+maxPathSum(root.Right))
}