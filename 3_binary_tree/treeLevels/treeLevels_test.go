package main

import "testing"

func TestTreeLevelsDFI_A(t *testing.T) {
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

	res := treeLevelsDFI(a)
	expected := [][]interface{}{
		{"a"}, {"b", "c"}, {"d", "e", "f"},
	}
	if !EqualNested(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestTreeLevelsBFI_A(t *testing.T) {
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

	res := treeLevelsBFI(a)
	expected := [][]interface{}{
		{"a"}, {"b", "c"}, {"d", "e", "f"},
	}
	if !EqualNested(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestTreeLevelsDFR_A(t *testing.T) {
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

	res := treeLevelsDFR(a)
	expected := [][]interface{}{
		{"a"}, {"b", "c"}, {"d", "e", "f"},
	}
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