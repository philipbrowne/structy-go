package main

import "testing"

func TestMaxPathSum_A(t *testing.T) {
	a := &Node{Val: 3}
	b := &Node{Val: 11}
	c := &Node{Val: 4}
	d := &Node{Val: 4}
	e := &Node{Val: -2}
	f := &Node{Val: 1}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f

	res := maxPathSum(a)
	if res != 18 {
		t.Errorf("expected %v, got %v", 18, res)
	}
}

func TestMaxPathSum_B(t *testing.T) {
	a := &Node{Val: 5}
	b := &Node{Val: 11}
	c := &Node{Val: 54}
	d := &Node{Val: 20}
	e := &Node{Val: 15}
	f := &Node{Val: 1}
	g := &Node{Val: 3}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	e.Left = f
	e.Right = g

	res := maxPathSum(a)
	if res != 59 {
		t.Errorf("expected %v, got %v", 59, res)
	}
}

func TestMaxPathSum_C(t *testing.T) {
	a := &Node{Val: -1}
	b := &Node{Val: -6}
	c := &Node{Val: -5}
	d := &Node{Val: -3}
	e := &Node{Val: 0}
	f := &Node{Val: -13}
	g := &Node{Val: -1}
	h := &Node{Val: -2}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	f.Right = h

	res := maxPathSum(a)
	if res != -8 {
		t.Errorf("expected %v, got %v", -8, res)
	}
}


func TestMaxPathSum_D(t *testing.T) {
	a := &Node{Val: 42}

	res := maxPathSum(a)
	if res != 42 {
		t.Errorf("expected %v, got %v", 42, res)
	}
}