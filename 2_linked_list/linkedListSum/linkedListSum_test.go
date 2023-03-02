package main

import "testing"

func TestSumIterative(t *testing.T) {
	a := &Node{Val: 2}
	b := &Node{Val: 8}
	c := &Node{Val: 3}
	d := &Node{Val: -1}
	e := &Node{Val: 7}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e

	res1 := linkedListSumIterative(a)
	if res1 != 19 {
		t.Errorf("expected 19, got %d", res1)
	}
	res2 := linkedListSumIterative(c)
	if res2 != 9 {
		t.Errorf("expected 9, got %d", res2)
	}
}

func TestSumRecursive(t *testing.T) {
	a := &Node{Val: 5}
	b := &Node{Val: 6}
	c := &Node{Val: 7}
	d := &Node{Val: -3}
	e := &Node{Val: 4}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e

	res1 := linkedListSumRecursive(a)
	if res1 != 19 {
		t.Errorf("expected 19, got %d", res1)
	}
	res2 := linkedListSumRecursive(c)
	if res2 != 8 {
		t.Errorf("expected 8, got %d", res2)
	}
}