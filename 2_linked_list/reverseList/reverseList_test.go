package main

import (
	"testing"
)

func TestReverseIterative(t *testing.T) {
	a := &Node{Val: "A"}
	b := &Node{Val: "B"}
	c := &Node{Val: "C"}
	d := &Node{Val: "D"}
	e := &Node{Val: "E"}
	f := &Node{Val: "F"}
	g := &Node{Val: "G"}
	h := &Node{Val: "H"}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	f.Next = g
	g.Next = h

	reverseListIteratively(a)
	if a.Next != nil {
		t.Errorf(`expected a.Next to be <nil>, resulted in %v`, a.Next)
	}
	if b.Next != a {
		t.Errorf(`expected b.Next to be %v, resulted in %v`, a, b.Next)
	}
	if g.Next != f {
		t.Errorf(`expected g.Next to be %v, resulted in %v`, f, g.Next)
	}
}

func TestReverseRecursive(t *testing.T) {
	a := &Node{Val: "A"}
	b := &Node{Val: "B"}
	c := &Node{Val: "C"}
	d := &Node{Val: "D"}
	e := &Node{Val: "E"}
	f := &Node{Val: "F"}
	g := &Node{Val: "G"}
	h := &Node{Val: "H"}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	f.Next = g
	g.Next = h

	reverseListRecursively(a)
	if a.Next != nil {
		t.Errorf(`expected a.Next to be <nil>, resulted in %v`, a.Next)
	}
	if b.Next != a {
		t.Errorf(`expected b.Next to be %v, resulted in %v`, a, b.Next)
	}
	if g.Next != f {
		t.Errorf(`expected g.Next to be %v, resulted in %v`, f, g.Next)
	}
}