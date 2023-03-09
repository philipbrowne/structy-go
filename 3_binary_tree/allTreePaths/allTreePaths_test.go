package main

import "testing"

func TestAllTreePaths_A(t *testing.T) {
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

	res := allTreePaths(a)
	expected := [][]interface{}{{"a", "b", "d"}, {"a", "b", "e"}, {"a", "c", "f"}}
	if !EqualNested(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestAllTreePaths_B(t *testing.T) {
	a := &Node{Val: "a"}
	b := &Node{Val: "b"}
	c := &Node{Val: "c"}
	d := &Node{Val: "d"}
	e := &Node{Val: "e"}
	f := &Node{Val: "f"}
	g := &Node{Val: "g"}
	h := &Node{Val: "h"}
	i := &Node{Val: "i"}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	e.Left = g
	e.Right = h
	f.Left = i

	res := allTreePaths(a)
	expected := [][]interface{}{{"a", "b", "d"}, {"a", "b", "e", "g"}, {"a", "b", "e", "h"}, {"a", "c", "f", "i"}}
	if !EqualNested(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestAllTreePaths_C(t *testing.T) {
	q := &Node{Val: "q"}
	r := &Node{Val: "r"}
	s := &Node{Val: "s"}
	tt := &Node{Val: "t"}
	u := &Node{Val: "u"}
	v := &Node{Val: "v"}
	

	q.Left = r
	q.Right = s
	r.Right = tt
	tt.Left = u
	u.Right = v

	res := allTreePaths(q)
	expected := [][]interface{}{{"q", "r", "t", "u", "v"}, {"q", "s"}}
	if !EqualNested(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestAllTreePaths_D(t *testing.T) {
	z := &Node{Val: "z"}
	
	res := allTreePaths(z)
	expected := [][]interface{}{{"z"}}
	if !EqualNested(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func EqualNested(a, b [][]interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !EqualInterface(v, b[i]) {
			return false
		}
	}
	return true
}

func EqualInterface(a, b []interface{}) bool {
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