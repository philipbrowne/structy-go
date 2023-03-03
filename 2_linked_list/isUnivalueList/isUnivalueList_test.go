package main

import "testing"

func TestUnivalueIterativeA(t *testing.T) {
	a := &Node{Val: 7}
	b := &Node{Val: 7}
	c := &Node{Val: 7}
	a.Next = b
	b.Next = c

	res := isUnivalueListIterative(a)
	if res != true {
		t.Errorf("expected result to be true, was %v", res)
	}
}

func TestUnivalueIterativeB(t *testing.T) {
	a := &Node{Val: 7}
	b := &Node{Val: 7}
	c := &Node{Val: 4}
	a.Next = b
	b.Next = c

	res := isUnivalueListIterative(a)
	if res != false {
		t.Errorf("expected result to be false, was %v", res)
	}
}

func TestUnivalueIterativeC(t *testing.T) {
	a := &Node{Val: 7}

	res := isUnivalueListIterative(a)
	if res != true {
		t.Errorf("expected result to be true, was %v", res)
	}
}

func TestUnivalueRecursiveA(t *testing.T) {
	a := &Node{Val: 7}
	b := &Node{Val: 7}
	c := &Node{Val: 7}
	a.Next = b
	b.Next = c

	res := isUnivalueListRecursive(a)
	if res != true {
		t.Errorf("expected result to be true, was %v", res)
	}
}

func TestUnivalueRecursiveB(t *testing.T) {
	a := &Node{Val: 7}
	b := &Node{Val: 7}
	c := &Node{Val: 4}
	a.Next = b
	b.Next = c

	res := isUnivalueListRecursive(a)
	if res != false {
		t.Errorf("expected result to be false, was %v", res)
	}
}

func TestUnivalueRecursiveC(t *testing.T) {
	a := &Node{Val: 7}

	res := isUnivalueListRecursive(a)
	if res != true {
		t.Errorf("expected result to be true, was %v", res)
	}
}