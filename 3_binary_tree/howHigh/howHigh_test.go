package main

import "testing"

func TestHowHigh_A(t *testing.T) {
	a := &Node{Val: "a"}
	b := &Node{Val: "b"}
	c := &Node{Val: "c"}
	d := &Node{Val: "d"}
	e := &Node{Val: "e"}
	f := &Node{Val: "f"}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f

	res := howHigh(a)
	if res != 2 {
		t.Errorf("expected %v, got %v", 2, res)
	}
}

func TestHowHigh_B(t *testing.T) {
	a := &Node{Val: "a"}
	b := &Node{Val: "b"}
	c := &Node{Val: "c"}
	d := &Node{Val: "d"}
	e := &Node{Val: "e"}
	f := &Node{Val: "f"}
	g := &Node{Val: "g"}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g

	res := howHigh(a)
	if res != 3 {
		t.Errorf("expected %v, got %v", 3, res)
	}
}

func TestHowHigh_C(t *testing.T) {
	a := &Node{Val: "a"}
	c := &Node{Val: "c"}

	a.Right = c

	res := howHigh(a)
	if res != 1 {
		t.Errorf("expected %v, got %v", 1, res)
	}
}

func TestHowHigh_D(t *testing.T) {
	a := &Node{Val: "a"}

	res := howHigh(a)
	if res != 0 {
		t.Errorf("expected %v, got %v", 0, res)
	}
}

func TestHowHigh_E(t *testing.T) {
	res := howHigh(nil)
	if res != -1 {
		t.Errorf("expected %v, got %v", -1, res)
	}
}