package main

type Node struct {
	Val interface{}
	Left *Node
	Right *Node
}

func main () {}

// Depth First Recursive
func allTreePaths(root *Node) [][]interface{} {
	 if root == nil {
		return [][]interface{}{}
	 }
	 if root.Left == nil && root.Right == nil {
		return [][]interface{}{{root.Val}}
	 }
	 paths := [][]interface{}{}
	 leftSubPaths := allTreePaths(root.Left)
	 for _, subPath := range leftSubPaths {
		toAppend := []interface{}{root.Val}
		toAppend = append(toAppend, subPath...)
		paths = append(paths, toAppend)
	 }
	 rightSubPaths := allTreePaths(root.Right)
	 for _, subPath := range rightSubPaths {
		toAppend := []interface{}{root.Val}
		toAppend = append(toAppend, subPath...)
		paths = append(paths, toAppend)
	 }
	 return paths
}