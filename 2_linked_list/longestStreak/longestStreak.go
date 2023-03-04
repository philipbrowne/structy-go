package main

type Node struct {
	Val interface{}
	Next *Node
}

func main () {}

// O(n) Runtime
func longestStreak(head *Node) int {
	currentStreak := 0
	maxStreak := 0
	currentNode := head
	var previousValue interface{}
	for currentNode != nil {
		if previousValue == nil || currentNode.Val == previousValue {
			currentStreak++
		} else {
			currentStreak = 1
		}
		previousValue = currentNode.Val
		if currentStreak > maxStreak {
			maxStreak = currentStreak
		}
		currentNode = currentNode.Next
	}
	return maxStreak
}