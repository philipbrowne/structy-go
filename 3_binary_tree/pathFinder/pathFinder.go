package main

type Node struct {
	Val interface{}
	Left *Node
	Right *Node
}

func main () {}

// Depth First Recursive - O(n^2) Runtime
func pathFinderA(root *Node, target interface{}) []interface{} {
	if root == nil {
		return nil
	}
	if root.Val == target {
		return []interface{}{root.Val}
	}
	valuesLeft := pathFinderA(root.Left, target)
	if valuesLeft != nil {
		toReturn  := []interface{}{root.Val}
		toReturn = append(toReturn, valuesLeft...)
		return toReturn
	}
	valuesRight := pathFinderA(root.Right, target)
	if valuesRight != nil {
		toReturn  := []interface{}{root.Val}
		toReturn = append(toReturn, valuesRight...)
		return toReturn
	}
	return nil
}

// More Optimal Depth First Recursive with Helper - O(n) Runtime
func pathFinderB(root *Node, target interface{}) []interface{} {
	result := pathFinderHelper(root, target)
	if result == nil {
		return nil
	} else {
		return Reverse(result)
	}
}

func pathFinderHelper(root *Node, target interface{}) []interface{} {
	if root == nil {
		return nil
	}
	if root.Val == target {
		return []interface{}{root.Val}
	}
	valuesLeft := pathFinderHelper(root.Left, target)
	if valuesLeft != nil {
		valuesLeft = append(valuesLeft, root.Val)
		return valuesLeft
	}
	valuesRight := pathFinderHelper(root.Right, target)
	if valuesRight != nil {
		valuesRight = append(valuesRight, root.Val)
		return valuesRight
	}
	return nil
}

func Reverse(input []interface{}) []interface{} {
    var output []interface{}

    for i := len(input) - 1; i >= 0; i-- {
        output = append(output, input[i])
    }

    return output
}