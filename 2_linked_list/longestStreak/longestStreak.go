package main

type Node struct {
	Val interface{}
	Next *Node
}

func main () {
	a := &Node{Val: 5}
	b := &Node{Val: 5}
	c := &Node{Val: 7}
	d := &Node{Val: 7}
	e := &Node{Val: 7}
	f := &Node{Val: 6}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
}

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