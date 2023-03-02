package main

import "testing"

func TestLinkedListFindIterative(t *testing.T) {
	a := &Node{Val: "A"}
	b := &Node{Val: "B"}
	c := &Node{Val: "C"}
	d := &Node{Val: "D"}
	a.Next = b
	b.Next = c
	c.Next = d

	res1 := linkedListFindIterative(a, "D")
	if res1 != true {
		t.Errorf("expected true, resulted in %v", res1)
	}
	res2 := linkedListFindIterative(a, "F")
	if res2 != false {
		t.Errorf("expected false, resulted in %v", res1)
	}
}

func TestLinkedListFindRecursive(t *testing.T) {
	a := &Node{Val: "A"}
	b := &Node{Val: "B"}
	c := &Node{Val: "C"}
	d := &Node{Val: "D"}
	a.Next = b
	b.Next = c
	c.Next = d

	res1 := linkedListFindRecursive(a, "D")
	if res1 != true {
		t.Errorf("expected true, resulted in %v", res1)
	}
	res2 := linkedListFindRecursive(b, "G")
	if res2 != false {
		t.Errorf("expected false, resulted in %v", res1)
	}
}