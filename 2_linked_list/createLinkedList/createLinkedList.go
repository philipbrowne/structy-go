package main

type Node struct {
	Val interface{}
	Next *Node
}

func main () {}

// Iterative Approach - O(n) Runtime
func createLinkedListIterative(values []interface{}) *Node {
	head := &Node{Val: nil}
	tail := head
	for _, val := range values {
		tail.Next = &Node{
			Val: val,
		}
		tail = tail.Next
	}
	return head.Next
}

// Recursive Approach - O(n) Runtime
func createLinkedListRecursive(values []interface{}, index ...int) *Node {
	if len(index) == 0 {
		index = append(index, 0)
	}
	if index[0] >= len(values) {
		return nil
	}
	head := &Node{
		Val: values[index[0]],
	}
	head.Next = createLinkedListRecursive(values, index[0] + 1)
	return head
}