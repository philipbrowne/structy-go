package main

// Binary Tree Node containing Integer Value
type Node struct {
	Val int
	Left *Node
	Right *Node
}

// Depth First Recursive - O(n) Runtime
func levelAverages(root *Node) []float64{
	levels := [][]int{}
	treeLevels(root, &levels, 0)
	averages := make([]float64, 0, len(levels))
	for _, level := range levels {
		averages = append(averages, average(level))
	}
	return averages
}

// Helper Recursive Function for Getting Values of Each Binary Tree Level
func treeLevels(root *Node, levels *[][]int, levelNum int) {
	if root == nil {
		return
	}
	if len(*levels) <= levelNum {
		*levels = append(*levels, []int{root.Val})
	} else {
		(*levels)[levelNum] = append((*levels)[levelNum], root.Val)
	}
	treeLevels(root.Left, levels, levelNum+1)
	treeLevels(root.Right, levels, levelNum+1)
}

// Utility Function
func average(nums []int) float64 {
	sum := 0.0
	for _, num := range nums {
		sum += float64(num)
	}
  return sum / float64(len(nums))
}