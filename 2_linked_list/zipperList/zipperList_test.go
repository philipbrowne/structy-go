package main

import "testing"

func TestZipperIterative(t *testing.T) {
	a := &Node{Val: "A"}
	b := &Node{Val: "B"}
	c := &Node{Val: "C"}
	d := &Node{Val: "D"}
	
	a.Next = b
	b.Next = c
	c.Next = d

	x := &Node{Val: "X"}
	y := &Node{Val: "Y"}
	z := &Node{Val: "Z"}

	x.Next = y
	y.Next = z

	zipperListIterative(a, x)
	if a.Next != x {
		t.Errorf(`Expected a.Next to be %v, got %v`, x, a.Next)
	}
	if x.Next != b {
		t.Errorf(`Expected x.Next to be %v, got %v`, b, x.Next)
	}
	if z.Next != d {
		t.Errorf(`Expected z.Next to be %v, got %v`, d, z.Next)
	}
}

func TestZipperRecursive(t *testing.T) {
	a := &Node{Val: "A"}
	b := &Node{Val: "B"}
	c := &Node{Val: "C"}
	d := &Node{Val: "D"}
	
	a.Next = b
	b.Next = c
	c.Next = d

	x := &Node{Val: "X"}
	y := &Node{Val: "Y"}
	z := &Node{Val: "Z"}

	x.Next = y
	y.Next = z

	zipperListRecursive(a, x)
	if a.Next != x {
		t.Errorf(`Expected a.Next to be %v, got %v`, x, a.Next)
	}
	if x.Next != b {
		t.Errorf(`Expected x.Next to be %v, got %v`, b, x.Next)
	}
	if z.Next != d {
		t.Errorf(`Expected z.Next to be %v, got %v`, d, z.Next)
	}
}