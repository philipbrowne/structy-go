package main

import "testing"

func TestValuesIteratively(t *testing.T) {
	a := &Node{Val: "A"}
	b := &Node{Val: "B"}
	c := &Node{Val: "C"}
	d := &Node{Val: "D"}
	a.Next = b
	b.Next = c
	c.Next = d
	res1 := linkedListValuesIterative(a)
	if !Equal(res1, []interface{}{"A", "B", "C", "D"}) {
		t.Errorf(`expected []interface{}{"A", "B", "C", "D"}, got %v`, res1)
	}
	res2 := linkedListValuesIterative(c)
	if !Equal(res2, []interface{}{"C", "D"}) {
		t.Errorf(`expected []interface{}{"C", "D"}, got %v`, res2)
	}
}

func TestValuesRecursively(t *testing.T) {
	a := &Node{Val: "A"}
	b := &Node{Val: "B"}
	c := &Node{Val: "C"}
	d := &Node{Val: "D"}
	a.Next = b
	b.Next = c
	c.Next = d
	res1 := linkedListValuesRecursive(a)
	if !Equal(res1, []interface{}{"A", "B", "C", "D"}) {
		t.Errorf(`expected []interface{}{"A", "B", "C", "D"}, got %v`, res1)
	}
	res2 := linkedListValuesRecursive(c)
	if !Equal(res2, []interface{}{"C", "D"}) {
		t.Errorf(`expected []interface{}{"C", "D"}, got %v`, res2)
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