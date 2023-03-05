package main

type Node struct {
	Val int
	Next *Node
}

func main () {

}

// Iterative Approach - O(max(n,m)) Runtime where n and m are the lengths of each list
func addListsIterative(head1, head2 *Node) *Node {
	head := &Node{}
	tail := head
	current1 := head1
	current2 := head2
	carry := 0
	for current1 != nil || current2 != nil || carry != 0 {
		num1 := 0
		if current1 != nil {
			num1 = current1.Val
		}
		num2 := 0
		if current2 != nil {
			num2 = current2.Val
		}
		sum := num1+num2+carry
		digit := sum%10
		if sum > 9 {
			carry = 1
		} else {
			carry = 0
		}
		tail.Next = &Node{Val: digit}
		tail = tail.Next
		if current1 != nil {
			current1 = current1.Next
		}
		if current2 != nil {
			current2 = current2.Next
		}
	}
	return head.Next
}

// Recursive Approach - O(max(n,m)) Runtime where n and m are the lengths of each list
func addListsRecursive(head1, head2 *Node, carry ...int) *Node {
	if len(carry) == 0 {
		carry = append(carry, 0)
	}
	if head1 == nil && head2 == nil && carry[0] == 0 {
		return nil
	}
	num1 := 0
	if head1 != nil {
		num1 = head1.Val
	}
	num2 := 0
	if head2 != nil {
		num2 = head2.Val
	}
	sum := num1 + num2 + carry[0]
	digit := sum%10
	if sum > 9 {
		carry[0] = 1
	} else {
		carry[0] = 0
	}
	result := &Node{Val: digit}
	var next1 *Node
	if head1 != nil {
		next1 = head1.Next
	}
	var next2 *Node
	if head2 != nil {
		next2 = head2.Next
	}
	result.Next = addListsRecursive(next1, next2, carry[0])
	return result
}