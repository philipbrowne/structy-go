package main

import "testing"

func TestInsertNodeIterativeA(t *testing.T) {
	a := &Node{Val: "A"}
	b := &Node{Val: "B"}
	c := &Node{Val: "C"}
	d := &Node{Val: "D"}
	a.Next = b
	b.Next = c
	c.Next = d
	insertNodeIterative(a, "X", 2)
	if b.Next.Val != "X" {
		t.Errorf(`expected X, got %v`, b.Next.Val)
	}
	if b.Next.Next != c {
		t.Errorf("expected %v, got %v", c, b.Next.Next)
	}
}

func TestInsertNodeIterativeB(t *testing.T) {
	a := &Node{Val: "A"}
	b := &Node{Val: "B"}
	c := &Node{Val: "C"}
	d := &Node{Val: "D"}
	a.Next = b
	b.Next = c
	c.Next = d
	insertNodeIterative(a, "V", 3)
	if c.Next.Val != "V" {
		t.Errorf(`expected V, got %v`, b.Next.Val)
	}
	if c.Next.Next != d {
		t.Errorf("expected %v, got %v", d, b.Next.Next)
	}
}

func TestInsertNodeIterativeC(t *testing.T) {
	a := &Node{Val: "A"}
	b := &Node{Val: "B"}
	c := &Node{Val: "C"}
	d := &Node{Val: "D"}
	a.Next = b
	b.Next = c
	c.Next = d
	insertNodeIterative(a, "M", 4)
	if d.Next == nil {
		t.Errorf("expected d.Next to be a node, got %v", d.Next)
	}
	if d.Next.Val != "M" {
		t.Errorf("expected %v, got %v", "M", d.Next.Val)
	}
}

func TestInsertNodeIterativeD(t *testing.T) {
	a := &Node{Val: "A"}
	b := &Node{Val: "B"}
	a.Next = b
	res := insertNodeIterative(a, "Z", 0)
	if res.Val != "Z" {
		t.Errorf("expected %v got %v", "Z", res.Val)
	}
	if res.Next != a {
		t.Errorf("expected %v, got %v", a, res.Next)
	}
}

func TestInsertNodeRecursiveA(t *testing.T) {
	a := &Node{Val: "A"}
	b := &Node{Val: "B"}
	c := &Node{Val: "C"}
	d := &Node{Val: "D"}
	a.Next = b
	b.Next = c
	c.Next = d
	insertNodeRecursive(a, "X", 2)
	if b.Next.Val != "X" {
		t.Errorf(`expected X, got %v`, b.Next.Val)
	}
	if b.Next.Next != c {
		t.Errorf("expected %v, got %v", c, b.Next.Next)
	}
}

func TestInsertNodeRecursiveB(t *testing.T) {
	a := &Node{Val: "A"}
	b := &Node{Val: "B"}
	c := &Node{Val: "C"}
	d := &Node{Val: "D"}
	a.Next = b
	b.Next = c
	c.Next = d
	insertNodeRecursive(a, "V", 3)
	if c.Next.Val != "V" {
		t.Errorf(`expected V, got %v`, b.Next.Val)
	}
	if c.Next.Next != d {
		t.Errorf("expected %v, got %v", d, b.Next.Next)
	}
}

func TestInsertNodeRecursiveC(t *testing.T) {
	a := &Node{Val: "A"}
	b := &Node{Val: "B"}
	c := &Node{Val: "C"}
	d := &Node{Val: "D"}
	a.Next = b
	b.Next = c
	c.Next = d
	insertNodeRecursive(a, "M", 4)
	if d.Next == nil {
		t.Errorf("expected d.Next to be a node, got %v", d.Next)
	}
	if d.Next.Val != "M" {
		t.Errorf("expected %v, got %v", "M", d.Next.Val)
	}
}

func TestInsertNodeRecursiveD(t *testing.T) {
	a := &Node{Val: "A"}
	b := &Node{Val: "B"}
	a.Next = b
	res := insertNodeRecursive(a, "Z", 0)
	if res.Val != "Z" {
		t.Errorf("expected %v got %v", "Z", res.Val)
	}
	if res.Next != a {
		t.Errorf("expected %v, got %v", a, res.Next)
	}
}