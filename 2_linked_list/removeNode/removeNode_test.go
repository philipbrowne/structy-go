package main

import (
	"testing"
)

func TestRemoveNodeIterativeA(t *testing.T) {
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

	removeNodeIterative(a, "C")
	if b.Next != d {
		t.Errorf("resulted in %v, expected value was %v", b.Next, d)
	}
}

func TestRemoveNodeIterativeB(t *testing.T) {
	x := &Node{Val: "X"}
	y := &Node{Val: "Y"}
	z := &Node{Val: "Z"}

	
	x.Next = y
	y.Next = z

	removeNodeIterative(x, "Z")
	if y.Next != nil {
		t.Errorf("resulted in %v, expected value was %v", y.Next, nil)
	}
}

func TestRemoveNodeIterativeC(t *testing.T) {
	q := &Node{Val: "Q"}
	r := &Node{Val: "R"}
	s := &Node{Val: "S"}

	
	q.Next = r
	r.Next = s

	res := removeNodeIterative(q, "Q")
	if res != r {
		t.Errorf("expected value was %v, got %v", r, res)
	}
}

func TestRemoveNodeIterativeD(t *testing.T) {
	n1 := &Node{Val: "H"}
	n2 := &Node{Val: "I"}
	n3 := &Node{Val: "J"}
	n4 := &Node{Val: "I"}
	
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	removeNodeIterative(n1, "I")
	
	if n1.Next != n3 {
		t.Errorf("expected %v, got %v", n3, n1.Next)
	}
}

func TestRemoveNodeIterativeE(t *testing.T) {
	T := &Node{Val: "T"}

	res := removeNodeIterative(T, "T")
	
	if res != nil {
		t.Errorf("expected %v, got %v", nil, res)
	}
}

func TestRemoveNodeRecursiveA(t *testing.T) {
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

	removeNodeRecursive(a, "C")
	if b.Next != d {
		t.Errorf("resulted in %v, expected value was %v", b.Next, d)
	}
}

func TestRemoveNodeRecursiveB(t *testing.T) {
	x := &Node{Val: "X"}
	y := &Node{Val: "Y"}
	z := &Node{Val: "Z"}

	
	x.Next = y
	y.Next = z

	removeNodeRecursive(x, "Z")
	if y.Next != nil {
		t.Errorf("resulted in %v, expected value was %v", y.Next, nil)
	}
}

func TestRemoveNodeRecursiveC(t *testing.T) {
	q := &Node{Val: "Q"}
	r := &Node{Val: "R"}
	s := &Node{Val: "S"}

	
	q.Next = r
	r.Next = s

	res := removeNodeRecursive(q, "Q")
	if res != r {
		t.Errorf("expected value was %v, got %v", r, res)
	}
}

func TestRemoveNodeRecursiveD(t *testing.T) {
	n1 := &Node{Val: "H"}
	n2 := &Node{Val: "I"}
	n3 := &Node{Val: "J"}
	n4 := &Node{Val: "I"}
	
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	removeNodeRecursive(n1, "I")
	
	if n1.Next != n3 {
		t.Errorf("expected %v, got %v", n3, n1.Next)
	}
}

func TestRemoveNodeRecursiveE(t *testing.T) {
	T := &Node{Val: "T"}

	res := removeNodeRecursive(T, "T")
	
	if res != nil {
		t.Errorf("expected %v, got %v", nil, res)
	}
}