package main

import "math"

type Node struct {
	Val interface{}
	Left *Node
	Right *Node
}

func main () {}

func howHigh(node *Node) float64 {
	if node == nil {
		return -1
	}
	return 1 + math.Max(howHigh(node.Left), howHigh(node.Right))
}