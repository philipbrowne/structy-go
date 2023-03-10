package main

type Node struct {
	Val interface{}
	Left *Node
	Right *Node
}

type LevelNode struct {
	Node *Node
	Level int
}

func main () {}

// Depth First Recursive - O(n) Runtime
func levelAverages(root *Node) []float64{
	levels := [][]interface{}{}
	treeLevelsHelper(root, &levels, 0)
	averages := []float64{}
	for _, level := range levels {
		averages = append(averages, average(level))
	}
	return averages
}

// Helper Function for Tree Levels Recursive
func treeLevelsHelper(root *Node, levels *[][]interface{}, levelNum int) {
	if root == nil {
		return
	}
	if len(*levels) <= levelNum {
		*levels = append(*levels, []interface{}{root.Val})
	} else {
		(*levels)[levelNum] = append((*levels)[levelNum], root.Val)
	}
	treeLevelsHelper(root.Left, levels, levelNum+1)
	treeLevelsHelper(root.Right, levels, levelNum+1)
}

func average(nums []interface{}) float64 {
	sum := 0.0
	for _, num := range nums {
		number := num.(int)
		sum += float64(number)
	}
	return sum / float64(len(nums))
}