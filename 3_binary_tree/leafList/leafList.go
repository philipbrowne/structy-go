package main

type Node struct {
	Val interface{}
	Left *Node
	Right *Node
}

func main () {}

func leafListDFI(root *Node) []interface{} {
	if root == nil {
		return []interface{}{}
	}
	leafs := []interface{}{}
	stack := []*Node{root}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
		if current.Left == nil && current.Right == nil {
			leafs = append(leafs, current.Val)
		}
	}
	return leafs
}

func leafListDFR(root *Node) []interface{} {
	leafs := &[]interface{}{}
	leafListHelper(root, leafs)
	return *leafs
}

func leafListHelper(root *Node, leafs *[]interface{}) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*leafs = append(*leafs, root.Val)
	}
	leafListHelper(root.Left, leafs)
	leafListHelper(root.Right, leafs)
}