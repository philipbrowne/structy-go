package main

import "testing"

func TestBottomRight_A(t *testing.T) {
	a := &Node{Val: 3}
	b := &Node{Val: 11}
	c := &Node{Val: 10}
	d := &Node{Val: 4}
	e := &Node{Val: -2}
	f := &Node{Val: 1}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f

	res := bottomRightValue(a)
	if res != 1 {
		t.Errorf("expected %v, got %v", 1, res)
	}
}

func TestBottomRight_B(t *testing.T) {
	a := &Node{Val: -1}
	b := &Node{Val: -6}
	c := &Node{Val: -5}
	d := &Node{Val: -3}
	e := &Node{Val: -4}
	f := &Node{Val: -13}
	g := &Node{Val: -2}
	h := &Node{Val: 6}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	e.Right= h

	res := bottomRightValue(a)
	if res != 6 {
		t.Errorf("expected %v, got %v", 6, res)
	}
}

func TestBottomRight_C(t *testing.T) {
	a := &Node{Val: -1}
	b := &Node{Val: -6}
	c := &Node{Val: -5}
	d := &Node{Val: -3}
	e := &Node{Val: -4}
	f := &Node{Val: -13}
	g := &Node{Val: -2}
	h := &Node{Val: 6}
	i := &Node{Val: 7}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	e.Right= h
	f.Left = i

	res := bottomRightValue(a)
	if res != 7 {
		t.Errorf("expected %v, got %v", 7, res)
	}
}

func TestBottomRight_D(t *testing.T) {
	a := &Node{Val: "a"}
	b := &Node{Val: "b"}
	c := &Node{Val: "c"}
	d := &Node{Val: "d"}
	e := &Node{Val: "e"}
	f := &Node{Val: "f"}


	a.Left = b
	a.Right = c
	b.Right = d
	d.Left = e
	e.Right = f

	res := bottomRightValue(a)
	if res != "f" {
		t.Errorf("expected %v, got %v", "f", res)
	}
}

func TestBottomRight_E(t *testing.T) {
	a := &Node{Val: 42}
	res := bottomRightValue(a)
	if res != 42 {
		t.Errorf("expected %v, got %v", 42, res)
	}
}