package main

import (
	"testing"
)

func TestTreeSumIterativeDF_A(t *testing.T) {
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
	res := treeSumIterativeDF(a)
	if res != 21 {
		t.Errorf("expected %v, got %v", 21, res)
	}
}

func TestTreeSumIterativeDF_B(t *testing.T) {
	a := &Node{Val: 1}
	b := &Node{Val: 6}
	c := &Node{Val: 0}
	d := &Node{Val: 3}
	e := &Node{Val: -6}
	f := &Node{Val: 2}
	g := &Node{Val: 2}
	h := &Node{Val: 2}
	
	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	f.Right = h
	res := treeSumIterativeDF(a)
	if res != 10 {
		t.Errorf("expected %v, got %v", 10, res)
	}
}

func TestTreeSumIterativeDF_C(t *testing.T) {
	res := treeSumIterativeDF(nil)
	if res != 0 {
		t.Errorf("expected %v, got %v", 0, res)
	}
}

func TestTreeSumRecursiveDF_A(t *testing.T) {
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
	res := treeSumRecursiveDF(a)
	if res != 21 {
		t.Errorf("expected %v, got %v", 21, res)
	}
}

func TestTreeSumRecursiveDF_B(t *testing.T) {
	a := &Node{Val: 1}
	b := &Node{Val: 6}
	c := &Node{Val: 0}
	d := &Node{Val: 3}
	e := &Node{Val: -6}
	f := &Node{Val: 2}
	g := &Node{Val: 2}
	h := &Node{Val: 2}
	
	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	f.Right = h
	res := treeSumRecursiveDF(a)
	if res != 10 {
		t.Errorf("expected %v, got %v", 10, res)
	}
}

func TestTreeSumRecursiveDF_C(t *testing.T) {
	res := treeSumRecursiveDF(nil)
	if res != 0 {
		t.Errorf("expected %v, got %v", 0, res)
	}
}

func TestTreeSumBF_A(t *testing.T) {
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
	res := treeSumBF(a)
	if res != 21 {
		t.Errorf("expected %v, got %v", 21, res)
	}
}

func TestTreeSumBF_B(t *testing.T) {
	a := &Node{Val: 1}
	b := &Node{Val: 6}
	c := &Node{Val: 0}
	d := &Node{Val: 3}
	e := &Node{Val: -6}
	f := &Node{Val: 2}
	g := &Node{Val: 2}
	h := &Node{Val: 2}
	
	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	f.Right = h
	res := treeSumBF(a)
	if res != 10 {
		t.Errorf("expected %v, got %v", 10, res)
	}
}

func TestTreeSumBF_C(t *testing.T) {
	res := treeSumBF(nil)
	if res != 0 {
		t.Errorf("expected %v, got %v", 0, res)
	}
}