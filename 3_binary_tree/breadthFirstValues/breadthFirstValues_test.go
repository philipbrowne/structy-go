package main

import (
	"testing"
)

func TestBreadthFirstValuesA(t *testing.T) {
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
	res := breadthFirstValues(a)
	if !Equal(res, []interface{}{"a", "b", "c", "d", "e", "f"}) {
		t.Errorf("expected %v, got %v", []interface{}{"a", "b", "c", "d", "e", "f"}, res)
	}
}

func TestBreadthFirstValuesB(t *testing.T) {
	a := &Node{Val: "a"}
	b := &Node{Val: "b"}
	c := &Node{Val: "c"}
	d := &Node{Val: "d"}
	e := &Node{Val: "e"}
	f := &Node{Val: "f"}
	g := &Node{Val: "g"}
	h := &Node{Val: "h"}
	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	f.Right = h
	res := breadthFirstValues(a)
	if !Equal(res, []interface{}{"a", "b", "c", "d", "e", "f", "g", "h"}) {
		t.Errorf("expected %v, got %v", []interface{}{"a", "b", "c", "d", "e", "f", "g", "h"}, res)
	}
}

func TestBreadthFirstValuesC(t *testing.T) {
	a := &Node{Val: "a"}
	res := breadthFirstValues(a)
	if !Equal(res, []interface{}{"a"}) {
		t.Errorf("expected %v, got %v", []interface{}{"a"}, res)
	}
}

func TestBreadthFirstValuesD(t *testing.T) {
	a := &Node{Val: "a"}
	b := &Node{Val: "b"}
	c := &Node{Val: "c"}
	d := &Node{Val: "d"}
	e := &Node{Val: "e"}
	x := &Node{Val: "x"}

	a.Right = b
	b.Left = c
	c.Left = x
	c.Right = d
	d.Right = e
	res := breadthFirstValues(a)
	if !Equal(res, []interface{}{"a", "b", "c", "x", "d", "e"}) {
		t.Errorf("expected %v, got %v", []interface{}{"a", "b", "c", "x", "d", "e"}, res)
	}
}

func TestBreadthFirstValuesE(t *testing.T) {
	res := breadthFirstValues(nil)
	if !Equal(res, []interface{}{}) {
		t.Errorf("expected %v, got %v", []interface{}{}, res)
	}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
