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

// Depth First Iterative - O(n) Runtime
func treeLevelsDFI(root *Node) [][]interface{} {
	if root == nil {
		return [][]interface{}{}
	}
	levels := [][]interface{}{}
	stack := []*LevelNode{{
			Node: root,
			Level: 0,
		}}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(levels) <= current.Level {
			levels = append(levels, []interface{}{current.Node.Val})
		} else {
			levels[current.Level] = append(levels[current.Level], current.Node.Val)
		}
		if current.Node.Right != nil {
			stack = append(stack, &LevelNode{
				Node: current.Node.Right,
				Level: current.Level+1,
			})
		}
		if current.Node.Left != nil {
			stack = append(stack, &LevelNode{
				Node: current.Node.Left,
				Level: current.Level+1,
			})
		}
	}
	return levels
}

// Breadth First Iterative - O(n) Runtime
func treeLevelsBFI(root *Node) [][]interface{} {
	if root == nil {
		return [][]interface{}{}
	}
	levels := [][]interface{}{}
	queue := []*LevelNode{{
		Node: root,
		Level: 0,
	}}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if len(levels) <= current.Level {
			levels = append(levels, []interface{}{current.Node.Val})
		} else {
			levels[current.Level] = append(levels[current.Level], current.Node.Val)
		}
		if current.Node.Left != nil {
			queue = append(queue, &LevelNode{
				Node: current.Node.Left,
				Level: current.Level+1,
			})
		}
		if current.Node.Right != nil {
			queue = append(queue, &LevelNode{
				Node: current.Node.Right,
				Level: current.Level+1,
			})
		}		
	}
	return levels	
}

// Depth First Recursive - O(n) Runtime
func treeLevelsDFR(root *Node) [][]interface{}{
	levels := [][]interface{}{}
	treeLevelsHelper(root, &levels, 0)
	return levels
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