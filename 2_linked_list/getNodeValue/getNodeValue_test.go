package main

import "testing"

func TestGetNodeValue(t *testing.T) {
	a := &Node{Val: "A"}
	b := &Node{Val: "B"}
	c := &Node{Val: "C"}
	d := &Node{Val: "D"}
	a.Next = b
	b.Next = c
	c.Next = d

	res1 := getNodeValueIterative(a, 2)
	if res1 != "C" {
		t.Errorf(`expected "C", got %s`, res1)
	}

	res2 := getNodeValueRecursive(a, 3)
	if res2 != "D" {
		t.Errorf(`expected "D", got %s`, res2)
	}
}