package main

import (
	"testing"
)

func TestDepthFirstValuesIterativeA(t *testing.T) {
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

	res := depthFirstValuesIterative(a)
	if !Equal(res, []interface{}{"a", "b", "d", "e", "c", "f"}) {
		t.Errorf("expected %v, got %v", []interface{}{"a", "b", "d", "e", "c", "f"}, res)
	}
}

func TestDepthFirstValuesIterativeB(t *testing.T) {
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

	res := depthFirstValuesIterative(a)
	if !Equal(res, []interface{}{"a", "b", "d", "e", "g", "c", "f"}) {
		t.Errorf("expected %v, got %v", []interface{}{"a", "b", "d", "e", "g", "c", "f"}, res)
	}
}

func TestDepthFirstValuesIterativeC(t *testing.T) {
	a := &Node{Val: "a"}
	
	res := depthFirstValuesIterative(a)
	if !Equal(res, []interface{}{"a"}) {
		t.Errorf("expected %v, got %v", []interface{}{"a"}, res)
	}
}

func TestDepthFirstValuesIterativeD(t *testing.T) {
	a := &Node{Val: "a"}
	b := &Node{Val: "b"}
	c := &Node{Val: "c"}
	d := &Node{Val: "d"}
	e := &Node{Val: "e"}
	
	a.Right = b
	b.Left = c
	c.Right = d
	d.Right = e

	res := depthFirstValuesIterative(a)
	if !Equal(res, []interface{}{"a", "b", "c", "d", "e"}) {
		t.Errorf("expected %v, got %v", []interface{}{"a", "b", "c", "d", "e"}, res)
	}
}

func TestDepthFirstValuesIterativeE(t *testing.T) {
	res := depthFirstValuesIterative(nil)
	if !Equal(res, []interface{}{}) {
		t.Errorf("expected %v, got %v", []interface{}{}, res)
	}
}

func TestDepthFirstValuesRecursiveA(t *testing.T) {
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

	res := depthFirstValuesRecursive(a)
	if !Equal(res, []interface{}{"a", "b", "d", "e", "c", "f"}) {
		t.Errorf("expected %v, got %v", []interface{}{"a", "b", "d", "e", "c", "f"}, res)
	}
}

func TestDepthFirstValuesRecursiveB(t *testing.T) {
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

	res := depthFirstValuesRecursive(a)
	if !Equal(res, []interface{}{"a", "b", "d", "e", "g", "c", "f"}) {
		t.Errorf("expected %v, got %v", []interface{}{"a", "b", "d", "e", "g", "c", "f"}, res)
	}
}

func TestDepthFirstValuesRecursiveC(t *testing.T) {
	a := &Node{Val: "a"}
	
	res := depthFirstValuesRecursive(a)
	if !Equal(res, []interface{}{"a"}) {
		t.Errorf("expected %v, got %v", []interface{}{"a"}, res)
	}
}

func TestDepthFirstValuesRecursiveD(t *testing.T) {
	a := &Node{Val: "a"}
	b := &Node{Val: "b"}
	c := &Node{Val: "c"}
	d := &Node{Val: "d"}
	e := &Node{Val: "e"}
	
	a.Right = b
	b.Left = c
	c.Right = d
	d.Right = e

	res := depthFirstValuesRecursive(a)
	if !Equal(res, []interface{}{"a", "b", "c", "d", "e"}) {
		t.Errorf("expected %v, got %v", []interface{}{"a", "b", "c", "d", "e"}, res)
	}
}

func TestDepthFirstValuesRecursiveE(t *testing.T) {
	res := depthFirstValuesRecursive(nil)
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
