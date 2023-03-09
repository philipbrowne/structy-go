package main

type Node struct {
	Val interface{}
	Next *Node
}

func main () {}

// Iterative - O(n) Runtime
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

func longestStreakRecursive(head *Node, streak int, longest int, prev ...interface{}) int {
	if len(prev) == 0 {
		prev = []interface{}{nil}
	}
	if streak > longest {
		longest = streak
	}
	if head == nil {
		return longest
	}
	if prev[0] == nil || head.Val == prev[0] {
		return longestStreakRecursive(head.Next, streak+1, longest, head.Val)
	} else {
		return longestStreakRecursive(head.Next, 1, longest, head.Val)
	}
}